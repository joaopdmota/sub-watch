package api

import (
	"boilerplate-go/internal/infra/http/webserver"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(httpService *webserver.HTTPService) {
	httpService.AddRoute(http.MethodGet, "/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})
}
