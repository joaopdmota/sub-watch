package httpclient

import (
	"net/http"
	"time"
)

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

type DefaultHTTPClient struct {
	baseUrl string
	client  *http.Client
	headers map[string]string
}

// ClientOption is a function that configures the HTTPClient
type ClientOption func(*DefaultHTTPClient)

// WithTimeout sets the timeout for HTTP requests
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *DefaultHTTPClient) {
		c.client.Timeout = timeout
	}
}

// WithHeaders sets custom headers for all requests
func WithHeaders(headers map[string]string) ClientOption {
	return func(c *DefaultHTTPClient) {
		c.headers = headers
	}
}

// WithTransport sets a custom transport for the HTTP client
func WithTransport(transport http.RoundTripper) ClientOption {
	return func(c *DefaultHTTPClient) {
		c.client.Transport = transport
	}
}

// WithHTTPClient allows using a completely custom http.Client
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *DefaultHTTPClient) {
		c.client = client
	}
}

// NewDefaultHTTPClient creates a new HTTP client with optional configuration
func NewDefaultHTTPClient(baseUrl string, opts ...ClientOption) *DefaultHTTPClient {
	client := &DefaultHTTPClient{
		baseUrl: baseUrl,
		client: &http.Client{
			Timeout: 30 * time.Second, // default timeout
		},
		headers: make(map[string]string),
	}

	// Apply all options
	for _, opt := range opts {
		opt(client)
	}

	return client
}

func (c *DefaultHTTPClient) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseUrl+url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}
