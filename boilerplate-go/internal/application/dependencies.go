package application

import (
	"boilerplate-go/internal/application/config"
	"boilerplate-go/internal/infra/http/webserver"
	"strconv"
)

func InitializeDependencies(envs *config.ConfigMap) *webserver.HTTPService {
	httpService := webserver.NewHTTPService(strconv.Itoa(envs.ApiPort), envs.ServiceName)

	return httpService
}
