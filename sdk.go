/*
  @author    Liran Polak
  @copyright Copyright (c) 2016, Spotinst
  @license   GPL-3.0
*/

package spotinstsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1"
	userAgent = "spotinst-sdk-go/" + libraryVersion
	defaultBaseProdURL = "http://dev.spotinst.com"
	defaultBaseTestURL = "http://dev.spotinst.com"
	defaultApiPort = 8081
	defaultOAuthPort = 9540
)

var (
	apiURL string
	oauthURL string
)

type Client struct {
	// This is our client structure.
	HttpClient   *http.Client

	// Spotinst makes a call to an authorization API using your username and
	// password, returning an 'Access Token' and a 'Refresh Token'.
	// Our use case does not require the refresh token, but we should implement
	// for completeness.
	AccessToken  string
	RefreshToken string
	Username     string
	Password     string
	ClientId     string
	ClientSecret string

	// Spotinst services.
	Group        *GroupService
}

type GroupResponse struct {
	Response struct {
				 Errors []Error `json:"errors"`
				 Items  []Group `json:"items"`
			 } `json:"response"`
}

type AuthResponse struct {
	Response struct {
				 Errors []Error `json:"errors"`
				 Items  []Token `json:"items"`
			 } `json:"response"`
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Error struct {
	Code    string `json:"code"`    // error code
	Message string `json:"message"` // human-readable message
	Field   string `json:"field"`
}

type ErrorResponseList struct {
	Response *http.Response // HTTP response that caused this error
	Errors   []Error
}

// NewClient returns a new Spotinst API client.
func NewClient(username, password, clientId, clientSecret string, test ...bool) (*Client, error) {
	var baseURL string

	if len(test) > 0 && test[0] {
		baseURL = defaultBaseTestURL
	} else {
		baseURL = defaultBaseProdURL
	}

	apiURL = fmt.Sprintf("%s:%d", baseURL, defaultApiPort)
	oauthURL = fmt.Sprintf("%s:%d", baseURL, defaultOAuthPort)

	accessToken, refreshToken, err := GetAuthTokens(username, password, clientId, clientSecret)

	if err != nil {
		return nil, err
	}

	c := &Client{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Username:     username,
		Password:     password,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		HttpClient:   &http.Client{},
	}

	c.Group = &GroupService{client: c}

	return c, nil
}

// GetAuthTokens creates an Authorization request to get an access and refresh token.
func GetAuthTokens(username, password, clientId, clientSecret string) (string, string, error) {
	res, err := http.PostForm(
		fmt.Sprintf("%s/token", oauthURL),
		url.Values{
			"grant_type":    {"password"},
			"username":      {username},
			"password":      {password},
			"client_id":     {clientId},
			"client_secret": {clientSecret},
		},
	)

	if err != nil {
		return "", "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	err = CheckResponse(res)
	if err != nil {
		return string(body), "", err
	}

	var authr AuthResponse
	err = json.Unmarshal(body, &authr)
	if err != nil {
		return string(body), "JSON Decode Error", err
	}

	var accessToken, refreshToken string
	for _, v := range authr.Response.Items {
		if v.AccessToken != "" {
			accessToken = v.AccessToken
		}

		if v.RefreshToken != "" {
			refreshToken = v.RefreshToken
		}
	}

	return accessToken, refreshToken, err
}

// NewRequest creates an API request.
// The path is expected to be a relative path and will be resolved
// according to the BaseURL of the Client. Paths should always be specified without a preceding slash.
func (client *Client) NewRequest(method, path string, payload interface{}) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s", apiURL, path)
	body := new(bytes.Buffer)
	if payload != nil {
		err := json.NewEncoder(body).Encode(payload)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.AccessToken))

	return req, nil
}

// Do sends an API request and returns the API response.
// The API response is JSON decoded and stored in the value pointed by v,
// or returned as an error if an API error has occurred.
// If v implements the io.Writer interface, the raw response body will be written to v,
// without attempting to decode it.
func (c *Client) Do(method, path string, payload, v interface{}) (*http.Response, error) {
	req, err := c.NewRequest(method, path, payload)
	if err != nil {
		return nil, err
	}

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = CheckResponse(res)

	if err != nil {
		return res, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, res.Body)
		} else {
			err = json.NewDecoder(res.Body).Decode(v)
		}
	}

	return res, err
}

func (c *Client) get(path string, v interface{}) (*http.Response, error) {
	return c.Do("GET", path, nil, v)
}

func (c *Client) post(path string, payload, v interface{}) (*http.Response, error) {
	return c.Do("POST", path, payload, v)
}

func (c *Client) put(path string, payload, v interface{}) (*http.Response, error) {
	return c.Do("PUT", path, payload, v)
}

func (c *Client) delete(path string, payload interface{}) (*http.Response, error) {
	return c.Do("DELETE", path, payload, nil)
}

// Error implements the error interface.
func (r *ErrorResponseList) Error() string {
	return fmt.Sprintf("%v %v: %d %s %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Errors[0].Code, r.Errors[0].Message)
}

// CheckResponse checks the API response for errors, and returns them if present.
// A response is considered an error if the status code is different than 2xx. Specific requests
// may have additional requirements, but this is sufficient in most of the cases.
func CheckResponse(res *http.Response) error {
	if code := res.StatusCode; 200 <= code && code <= 299 {
		return nil
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var r GroupResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return err
	}

	return &ErrorResponseList{Response: res, Errors: r.Response.Errors}
}
