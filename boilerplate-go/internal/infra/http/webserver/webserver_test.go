package webserver_test

import (
	"boilerplate-go/internal/infra/http/webserver"
	"io"
	"net/http"
	"net/http/httptest" // Corrected import path
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewHTTPService(t *testing.T) {
	port := "8080"
	service := webserver.NewHTTPService(port, "test")

	assert.NotNil(t, service)
}

func TestAddRoute(t *testing.T) {
	port := "8080"
	service := webserver.NewHTTPService(port, "test")

	getHandler := func(c echo.Context) error {
		return c.String(http.StatusOK, "GET OK")
	}
	postHandler := func(c echo.Context) error {
		return c.String(http.StatusOK, "POST OK")
	}
	putHandler := func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT OK")
	}
	deleteHandler := func(c echo.Context) error {
		return c.String(http.StatusOK, "DELETE OK")
	}

	service.AddRoute(http.MethodGet, "/test/get", getHandler)
	service.AddRoute(http.MethodPost, "/test/post", postHandler)
	service.AddRoute(http.MethodPut, "/test/put", putHandler)
	service.AddRoute(http.MethodDelete, "/test/delete", deleteHandler)

	ts := httptest.NewServer(service.Echo())
	defer ts.Close()

	tests := []struct {
		method       string
		url          string
		expectedBody string
	}{
		{http.MethodGet, "/test/get", "GET OK"},
		{http.MethodPost, "/test/post", "POST OK"},
		{http.MethodPut, "/test/put", "PUT OK"},
		{http.MethodDelete, "/test/delete", "DELETE OK"},
	}

	for _, tt := range tests {
		t.Run(tt.method, func(t *testing.T) {
			var resp *http.Response
			var err error

			switch tt.method {
			case http.MethodGet:
				resp, err = http.Get(ts.URL + tt.url)
			case http.MethodPost:
				resp, err = http.Post(ts.URL+tt.url, "application/json", nil)
			case http.MethodPut:
				req, _ := http.NewRequest(http.MethodPut, ts.URL+tt.url, nil)
				resp, err = http.DefaultClient.Do(req)
			case http.MethodDelete:
				req, _ := http.NewRequest(http.MethodDelete, ts.URL+tt.url, nil)
				resp, err = http.DefaultClient.Do(req)
			}

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			body, err := io.ReadAll(resp.Body)
			assert.NoError(t, err)
			resp.Body.Close()

			assert.Equal(t, tt.expectedBody, string(body))
		})
	}
}