package application

import (
	"net/http"
	"strconv"
	"sub-watch/application/config"
	"sub-watch/infra/http/router"

	"github.com/labstack/echo/v4"
)

func InitializeDependencies(envs *config.ConfigMap) *router.HTTPService {
	httpService := router.NewHTTPService(strconv.Itoa(envs.ApiPort))

	httpService.AddRoute(http.MethodGet, "/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	return httpService
}
