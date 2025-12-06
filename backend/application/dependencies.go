package application

import (
	"net/http"
	httpclient "sub-watch/infra/http/client"
	"time"
)

func InitializeDependencies() {
	httpclient.NewDefaultHTTPClient(&http.Client{
		Timeout: 5 * time.Second,
	})
}
