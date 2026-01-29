package http_infra

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(httpService *HTTPService) {
	httpService.AddRoute(http.MethodGet, "/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})
}
