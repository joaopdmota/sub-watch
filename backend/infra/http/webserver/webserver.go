package webserver

import (
	"log"

	router "sub-watch/infra/http/router"

	"github.com/labstack/echo/v4"
)

type WebServer struct {
	Port    string
	Service *router.HTTPService
}

func NewWebServer(port string) *WebServer {
	return &WebServer{
		Port:    port,
		Service: router.NewHTTPService(port),
	}
}

func (ws *WebServer) AddRoute(method, path string, handler echo.HandlerFunc) {
	ws.Service.AddRoute(method, path, handler)
}

func (ws *WebServer) Start() {
	if err := ws.Service.Start(); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}