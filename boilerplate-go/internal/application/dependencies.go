package application

import (
	"boilerplate-go/internal/application/config"
	http_infra "boilerplate-go/internal/infra/http"
	"strconv"
)

func InitializeDependencies(envs *config.ConfigMap) *http_infra.HTTPService {
	httpService := http_infra.NewHTTPService(strconv.Itoa(envs.ApiPort), envs.ServiceName)

	return httpService
}
