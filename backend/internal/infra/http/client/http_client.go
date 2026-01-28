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

type ClientOption func(*DefaultHTTPClient)

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *DefaultHTTPClient) {
		c.client.Timeout = timeout
	}
}

func WithHeaders(headers map[string]string) ClientOption {
	return func(c *DefaultHTTPClient) {
		c.headers = headers
	}
}

func WithTransport(transport http.RoundTripper) ClientOption {
	return func(c *DefaultHTTPClient) {
		c.client.Transport = transport
	}
}

func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *DefaultHTTPClient) {
		c.client = client
	}
}

func NewDefaultHTTPClient(baseUrl string, opts ...ClientOption) *DefaultHTTPClient {
	client := &DefaultHTTPClient{
		baseUrl: baseUrl,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		headers: make(map[string]string),
	}

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
