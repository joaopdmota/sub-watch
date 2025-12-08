package application

import (
	"boilerplate-go/application/config"
	"boilerplate-go/infra/http/webserver"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InitializeDependencies(envs *config.ConfigMap) *webserver.HTTPService {
	httpService := webserver.NewHTTPService(strconv.Itoa(envs.ApiPort), envs.ServiceName)

	httpService.AddRoute(http.MethodGet, "/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	return httpService
}
