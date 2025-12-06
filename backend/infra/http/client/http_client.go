package httpclient

import "net/http"

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

type DefaultHTTPClient struct {
	client *http.Client
}

func NewDefaultHTTPClient(client *http.Client) *DefaultHTTPClient {
	return &DefaultHTTPClient{client: client}
}

func (c *DefaultHTTPClient) Get(url string) (*http.Response, error) {
	return c.client.Get(url)
}
