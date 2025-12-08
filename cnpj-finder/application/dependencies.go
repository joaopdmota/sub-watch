package application

import (
	"cnpj-finder/application/config"
	"cnpj-finder/application/services"
	"cnpj-finder/application/usecases"
	httpclient "cnpj-finder/infra/http/client"
	"cnpj-finder/infra/http/handlers"
	"cnpj-finder/infra/http/webserver"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func InitializeDependencies(envs *config.ConfigMap) *webserver.HTTPService {
	cnpjFinderTimeout := time.Duration(envs.CnpjFinderClientTimeout) * time.Millisecond

	httpClient := httpclient.NewDefaultHTTPClient(envs.CnpjFinderClientUrl, httpclient.WithTimeout(cnpjFinderTimeout))

	receitaCNPJService := services.NewReceitaCNPJService(httpClient)

	cnpjUseCase := usecases.NewGetCNPJUseCase(*receitaCNPJService)

	documentHandler := handlers.NewDocumentHandler(cnpjUseCase)

	httpService := webserver.NewHTTPService(strconv.Itoa(envs.ApiPort), envs.ServiceName)

	httpService.AddRoute(http.MethodGet, "/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	httpService.AddRoute(http.MethodGet, "/consulta/:document", documentHandler.Get)

	return httpService
}
