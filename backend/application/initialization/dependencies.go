package app_init

import (
	"net/http"
	"strconv"
	"sub-watch-backend/application/auth"
	"sub-watch-backend/application/config"
	"sub-watch-backend/application/usecases"
	"sub-watch-backend/infra/database"
	"sub-watch-backend/infra/database/adapters"
	"sub-watch-backend/infra/http/handlers"
	"sub-watch-backend/infra/http/webserver"
	"sub-watch-backend/infra/repositories"
	"sub-watch-backend/infra/validator"
	"sub-watch-backend/pkg/date"
	"sub-watch-backend/pkg/hash"
	id "sub-watch-backend/pkg/uuid"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "sub-watch-backend/docs"
)

func InitializeDependencies(envs *config.ConfigMap) *webserver.HTTPService {
	httpService := webserver.NewHTTPService(strconv.Itoa(envs.ApiPort), envs.ServiceName)

	httpService.AddRoute(http.MethodGet, "/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	httpService.AddRoute(http.MethodGet, "/swagger/*", echoSwagger.WrapHandler)

	dbAdapter := adapters.NewPostgresAdapter()
	dbInstance, err := database.NewConnection(dbAdapter)
	if err != nil {
		panic(err)
	}

	idProvider := id.NewUUIDProvider()
	hasher := hash.NewBcryptHasher()
	dateProvider := date.NewDateProvider()
	validatorAdapter := validator.NewGoPlaygroundValidator()

	_ = auth.NewJWTService(envs)

	userRepo := repositories.NewUserRepository(dbInstance)
	categoryRepo := repositories.NewCategoryRepository(dbInstance)

	authLoginUseCase := usecases.NewAuthLoginUseCase(userRepo, hasher)

	getUserUseCase := usecases.NewGetUserUseCase(userRepo)
	listUsersUseCase := usecases.NewListUsersUseCase(userRepo)
	createUserUseCase := usecases.NewCreateUserUseCase(userRepo, idProvider, hasher, dateProvider, validatorAdapter)

	getCategoryUseCase := usecases.NewGetCategoryUseCase(categoryRepo)
	listCategoriesUseCase := usecases.NewListCategoriesUseCase(categoryRepo)

	userHandler := handlers.NewUserHandler(listUsersUseCase, getUserUseCase, createUserUseCase)
	categoryHandler := handlers.NewCategoryHandler(getCategoryUseCase, listCategoriesUseCase)

	authHandler := handlers.NewAuthHandler(authLoginUseCase)

	prefix := "api/"

	categoriesGroup := httpService.Group(prefix + "categories")

	categoriesGroup.GET("/:id", categoryHandler.GetCategory)
	categoriesGroup.GET("", categoryHandler.ListCategories)

	usersGroup := httpService.Group(prefix + "users")

	usersGroup.GET("/:id", userHandler.GetUser)
	usersGroup.GET("", userHandler.ListUsers)
	usersGroup.POST("", userHandler.CreateUser)	

	authGroup := httpService.Group(prefix + "auth")
	authGroup.POST("/login", authHandler.Login)

	return httpService
}
