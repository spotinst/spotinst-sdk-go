package spotinst

import (
	"fmt"
	"net/http"
)

const (
	apiURL    = "https://api.spotinst.io"
	oauthURL  = "https://oauth.spotinst.io"
	mediaType = "application/json"
	userAgent = SDKName + "/" + SDKVersion
)

type Credentials struct {
	Email        string `json:"username"`
	Password     string `json:"password"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Token        string `json:"token"`
}

type Response struct {
	Response struct {
		Errors []Error       `json:"errors"`
		Items  []interface{} `json:"items"`
	} `json:"response"`
}

type Error struct {
	// Error code
	Code string `json:"code"`

	// Human-readable message
	Message string `json:"message"`

	// The field in error
	Field string `json:"field"`
}

// An ErrorResponse reports the errors caused by an API request
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// A list of errors
	Errors []Error
}

// Error implements the error interface.
func (r *ErrorResponse) Error() string {
	if len(r.Errors) > 0 {
		return fmt.Sprintf("%v %v: %d %s %v",
			r.Response.Request.Method, r.Response.Request.URL,
			r.Response.StatusCode, r.Errors[0].Code, r.Errors[0].Message)
	} else {
		return fmt.Sprintf("%v %v: %d",
			r.Response.Request.Method, r.Response.Request.URL,
			r.Response.StatusCode)
	}
}
