package http_infra

import (
	"boilerplate-go/internal/infra/http/middlewares"
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

type HTTPService struct {
	server *http.Server
	echo   *echo.Echo
}

func NewHTTPService(port string, serviceName string) *HTTPService {
    e := echo.New()

    e.Use(middlewares.Logger())
    e.Use(middlewares.RequestID())
    e.Use(otelecho.Middleware(serviceName))

    server := &http.Server{
        Addr:    ":" + port,
        Handler: e,
    }

    return &HTTPService{
        server: server,
        echo:   e,
    }
}

func (s *HTTPService) AddRoute(method, pattern string, handler echo.HandlerFunc) {
	switch method {
	case http.MethodGet:
		s.echo.GET(pattern, handler)
	case http.MethodPost:
		s.echo.POST(pattern, handler)
	case http.MethodPut:
		s.echo.PUT(pattern, handler)
	case http.MethodDelete:
		s.echo.DELETE(pattern, handler)
	default:
		s.echo.Any(pattern, func(c echo.Context) error {
			return echo.NewHTTPError(http.StatusMethodNotAllowed, fmt.Sprintf("Método %s não é suportado", method))
		})
	}
}

func (s *HTTPService) Start() error {
	fmt.Printf("Servidor rodando na porta %s...\n", s.server.Addr)
	return s.echo.StartServer(s.server)
}

func (s *HTTPService) Stop(ctx context.Context) error {
	fmt.Println("Parando o servidor...")
	return s.server.Shutdown(ctx)
}

func (s *HTTPService) Echo() *echo.Echo {
	return s.echo
}
