package application

import (
	"net/http"
	"strconv"
	"sub-watch/application/config"
	"sub-watch/application/services"
	"sub-watch/application/usecases"
	"sub-watch/infra/database"
	httpclient "sub-watch/infra/http/client"
	"sub-watch/infra/http/handlers"
	"sub-watch/infra/http/router"
	"sub-watch/infra/repositories"
	"time"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitializeDependencies(envs *config.ConfigMap) *router.HTTPService {
	httpclient.NewDefaultHTTPClient(&http.Client{
		Timeout: 5 * time.Second,
	})

	db := database.NewPostgresAdapter()

	pool, err := database.NewConnection(db)
	if err != nil {
		panic(err)
	}

	userRepository := repositories.NewUserRepository(pool)
	userService := services.NewUserService(userRepository)
	
	listUsersUseCase := usecases.NewListUsersUseCase(userService)
	getUserUseCase := usecases.NewGetUserUseCase(userService)

	handler := handlers.NewUserHandler(listUsersUseCase, getUserUseCase)

	httpService := router.NewHTTPService(strconv.Itoa(envs.ApiPort))

	httpService.AddRoute(http.MethodGet, "/swagger/*", echoSwagger.WrapHandler)
	httpService.AddRoute(http.MethodGet, "/users/:id", handler.GetUser)
	httpService.AddRoute(http.MethodGet, "/users", handler.ListUsers)

	return httpService
}
