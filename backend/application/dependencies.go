package application

import (
	"net/http"
	"strconv"
	"sub-watch-backend/application/config"
	"sub-watch-backend/application/usecases"
	"sub-watch-backend/infra/database"
	"sub-watch-backend/infra/database/adapters"
	"sub-watch-backend/infra/http/handlers"
	"sub-watch-backend/infra/http/webserver"
	"sub-watch-backend/infra/repositories"
	"sub-watch-backend/pkg/date"
	"sub-watch-backend/pkg/hash"
	id "sub-watch-backend/pkg/uuid"

	"github.com/labstack/echo/v4"
)

func InitializeDependencies(envs *config.ConfigMap) *webserver.HTTPService {
	httpService := webserver.NewHTTPService(strconv.Itoa(envs.ApiPort), envs.ServiceName)

	httpService.AddRoute(http.MethodGet, "/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	dbAdapter := adapters.NewPostgresAdapter()
	dbInstance, err := database.NewConnection(dbAdapter)
	if err != nil {
		panic(err)
	}

	idProvider := id.NewUUIDProvider()
	hasher := hash.NewBcryptHasher()
	dateProvider := date.NewDateProvider()

	userRepo := repositories.NewUserRepository(dbInstance)
	getUserUseCase := usecases.NewGetUserUseCase(userRepo)
	listUsersUseCase := usecases.NewListUsersUseCase(userRepo)
	createUserUseCase := usecases.NewCreateUserUseCase(userRepo, idProvider, hasher, dateProvider)

	userHandler := handlers.NewUserHandler(listUsersUseCase, getUserUseCase, createUserUseCase)

	httpService.AddRoute(http.MethodGet, "/users", userHandler.ListUsers)
	httpService.AddRoute(http.MethodGet, "/users/:id", userHandler.GetUser)
	httpService.AddRoute(http.MethodPost, "/users", userHandler.CreateUser)

	return httpService
}
