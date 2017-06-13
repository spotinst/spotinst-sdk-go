package spotinst

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// DefaultTransport returns a new http.Transport with the same default values
// as http.DefaultTransport, but with idle connections and KeepAlives disabled.
func defaultTransport() *http.Transport {
	transport := defaultPooledTransport()
	transport.DisableKeepAlives = true
	transport.MaxIdleConnsPerHost = -1
	return transport
}

// DefaultPooledTransport returns a new http.Transport with similar default
// values to http.DefaultTransport. Do not use this for transient transports as
// it can leak file descriptors over time. Only use this for transports that
// will be re-used for the same host(s).
func defaultPooledTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
		DisableKeepAlives:   false,
		MaxIdleConnsPerHost: 1,
	}
}

// defaultHttpClient returns a new http.Client with similar default values to
// http.Client, but with a non-shared Transport, idle connections disabled, and
// KeepAlives disabled.
func defaultHttpClient() *http.Client {
	return &http.Client{
		Transport: defaultTransport(),
	}
}

// defaultHttpPooledClient returns a new http.Client with the same default values
// as http.Client, but with a shared Transport. Do not use this function
// for transient clients as it can leak file descriptors over time. Only use
// this for clients that will be re-used for the same host(s).
func defaultHttpPooledClient() *http.Client {
	return &http.Client{
		Transport: defaultPooledTransport(),
	}
}

// DefaultConfig returns a default configuration for the client. By default this
// will pool and reuse idle connections to API. If you have a long-lived
// client object, this is the desired behavior and should make the most efficient
// use of the connections to API. If you don't reuse a client object , which
// is not recommended, then you may notice idle connections building up over
// time. To avoid this, use the DefaultNonPooledConfig() instead.
func defaultPooledConfig() *clientConfig {
	return defaultConfig(defaultPooledTransport)
}

// DefaultNonPooledConfig returns a default configuration for the client which
// does not pool connections. This isn't a recommended configuration because it
// will reconnect to API on every request, but this is useful to avoid the
// accumulation of idle connections if you make many client objects during the
// lifetime of your application.
func defaultNonPooledConfig() *clientConfig {
	return defaultConfig(defaultTransport)
}

// defaultConfig returns the default configuration for the client, using the
// given function to make the transport.
func defaultConfig(transportFn func() *http.Transport) *clientConfig {
	return &clientConfig{
		address:     DefaultAddress,
		scheme:      DefaultScheme,
		userAgent:   DefaultUserAgent,
		contentType: DefaultContentType,
		httpClient: &http.Client{
			Transport: transportFn(),
		},
	}
}

// Client provides a client to the API.
type Client struct {
	config              *clientConfig
	AwsGroupService     AwsGroupService
	DeploymentService   DeploymentService
	CertificateService  CertificateService
	BalancerService     BalancerService
	HealthCheckService  HealthCheckService
	SubscriptionService SubscriptionService
}

// NewClient returns a new client.
func NewClient(opts ...ClientOptionFunc) (*Client, error) {
	config := defaultPooledConfig()

	for _, o := range opts {
		o(config)
	}

	client := &Client{config: config}

	// Elastigroup services.
	client.AwsGroupService = &AwsGroupServiceOp{client}

	// Load Balancer services.
	client.DeploymentService = &DeploymentServiceOp{client}
	client.CertificateService = &CertificateServiceOp{client}
	client.BalancerService = &BalancerServiceOp{client}

	// Health Check services.
	client.HealthCheckService = &HealthCheckServiceOp{client}

	// Subscription services.
	client.SubscriptionService = &SubscriptionServiceOp{client}

	return client, nil
}

// newRequest is used to create a new request.
func (c *Client) newRequest(method, path string) *request {
	req := &request{
		config: c.config,
		method: method,
		url: &url.URL{
			Scheme: c.config.scheme,
			Host:   c.config.address,
			Path:   path,
		},
		params: make(map[string][]string),
		header: make(http.Header),
	}
	if token := c.config.token; token != "" {
		req.header.Set("Authorization", "Bearer "+token)
	}
	if accountID := c.config.accountID; accountID != "" {
		req.params.Set("accountId", accountID)
	}
	return req
}

// doRequest runs a request with our client.
func (c *Client) doRequest(r *request) (time.Duration, *http.Response, error) {
	req, err := r.toHTTP()
	if err != nil {
		return 0, nil, err
	}
	c.dumpRequest(req)
	start := time.Now()
	resp, err := c.config.httpClient.Do(req)
	diff := time.Now().Sub(start)
	c.dumpResponse(resp)
	return diff, resp, err
}

// errorf logs to the error log.
func (c *Client) errorf(format string, args ...interface{}) {
	if c.config.errorlog != nil {
		c.config.errorlog.Printf(format, args...)
	}
}

// infof logs informational messages.
func (c *Client) infof(format string, args ...interface{}) {
	if c.config.infolog != nil {
		c.config.infolog.Printf(format, args...)
	}
}

// tracef logs to the trace log.
func (c *Client) tracef(format string, args ...interface{}) {
	if c.config.tracelog != nil {
		c.config.tracelog.Printf(format, args...)
	}
}

// dumpRequest dumps the given HTTP request to the trace log.
func (c *Client) dumpRequest(r *http.Request) {
	if c.config.tracelog != nil && r != nil {
		out, err := httputil.DumpRequestOut(r, true)
		if err == nil {
			c.tracef("%s\n", string(out))
		}
	}
}

// dumpResponse dumps the given HTTP response to the trace log.
func (c *Client) dumpResponse(resp *http.Response) {
	if c.config.tracelog != nil && resp != nil {
		out, err := httputil.DumpResponse(resp, true)
		if err == nil {
			c.tracef("%s\n", string(out))
		}
	}
}
