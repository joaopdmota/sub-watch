package httpclient_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	httpclient "sub-watch/infra/http/client"

	"github.com/stretchr/testify/assert"
)

func TestDefaultHTTPClient_Get_Success(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer mockServer.Close()

	client := httpclient.NewDefaultHTTPClient(mockServer.Client())

	resp, err := client.Get(mockServer.URL)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDefaultHTTPClient_Get_Error(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer mockServer.Close()

	client := httpclient.NewDefaultHTTPClient(mockServer.Client())

	resp, err := client.Get(mockServer.URL)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
