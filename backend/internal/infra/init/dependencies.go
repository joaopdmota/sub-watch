package infra_init

import (
	"log"
	"net/http"
	"strconv"
	"sub-watch-backend/internal/application"
	"sub-watch-backend/internal/application/auth"
	"sub-watch-backend/internal/application/config"
	"sub-watch-backend/internal/application/usecases"
	"sub-watch-backend/internal/infra/database"
	"sub-watch-backend/internal/infra/database/adapters"
	"sub-watch-backend/internal/infra/http/handlers"
	"sub-watch-backend/internal/infra/http/webserver"
	"sub-watch-backend/internal/infra/repositories"
	"sub-watch-backend/internal/infra/validator"
	"sub-watch-backend/internal/pkg/date"
	"sub-watch-backend/internal/pkg/hash"
	id "sub-watch-backend/internal/pkg/uuid"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "sub-watch-backend/docs"
)

type Platform struct {
	Envs      *config.ConfigMap
	Log       application.Logger
	DB        database.Database
	ID        id.UuidProvider
	Hasher    hash.PasswordHasher
	Date      date.DateProvider
	Validator application.Validator
}

func InitializeDependencies(envs *config.ConfigMap, logger application.Logger) *webserver.HTTPService {
	dbAdapter := adapters.NewPostgresAdapter()
	dbInstance, err := database.NewConnection(dbAdapter)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	platform := &Platform{
		Envs:      envs,
		Log:       logger,
		DB:        dbInstance,
		ID:        id.NewUUIDProvider(),
		Hasher:    hash.NewBcryptHasher(),
		Date:      date.NewDateProvider(),
		Validator: validator.NewGoPlaygroundValidator(),
	}

	httpService := webserver.NewHTTPService(strconv.Itoa(envs.ApiPort), envs.ServiceName)

	httpService.AddRoute(http.MethodGet, "/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})
	httpService.AddRoute(http.MethodGet, "/swagger/*", echoSwagger.WrapHandler)

	initUserFeature(httpService, platform)
	initAuthFeature(httpService, platform)
	initSubscriptionFeature(httpService, platform)
	initPaymentMethodFeature(httpService, platform)

	return httpService
}

func initUserFeature(svc *webserver.HTTPService, p *Platform) {
	repo := repositories.NewUserRepository(p.DB)
	
	listUC := usecases.NewListUsersUseCase(repo)
	getUC := usecases.NewGetUserUseCase(repo)
	createUC := usecases.NewCreateUserUseCase(repo, p.ID, p.Hasher, p.Date, p.Validator, p.Log)
	
	handler := handlers.NewUserHandler(listUC, getUC, createUC)

	group := svc.Group("api/users")
	group.GET("", handler.ListUsers)
	group.GET("/:id", handler.GetUser)
	group.POST("", handler.CreateUser)
}

func initAuthFeature(svc *webserver.HTTPService, p *Platform) {
	_ = auth.NewJWTService(p.Envs)
	repo := repositories.NewUserRepository(p.DB)
	
	loginUC := usecases.NewAuthLoginUseCase(repo, p.Hasher)
	handler := handlers.NewAuthHandler(loginUC)

	group := svc.Group("api/auth")
	group.POST("/login", handler.Login)
}

func initSubscriptionFeature(svc *webserver.HTTPService, p *Platform) {
	repo := repositories.NewSubscriptionRepository(p.DB)
	
	createUC := usecases.NewCreateSubscriptionUseCase(repo, p.ID, p.Date, p.Validator, p.Log)
	listUC := usecases.NewListSubscriptionsUseCase(repo)
	getUC := usecases.NewGetSubscriptionUseCase(repo)
	updateUC := usecases.NewUpdateSubscriptionUseCase(repo, p.Date, p.Validator, p.Log)
	deleteUC := usecases.NewDeleteSubscriptionUseCase(repo, p.Log)

	handler := handlers.NewSubscriptionHandler(createUC, listUC, getUC, updateUC, deleteUC)

	group := svc.Group("api/subscriptions")
	group.POST("", handler.CreateSubscription)
	group.GET("", handler.ListSubscriptions)
	group.GET("/:id", handler.GetSubscription)
	group.PUT("/:id", handler.UpdateSubscription)
	group.DELETE("/:id", handler.DeleteSubscription)
}

func initPaymentMethodFeature(svc *webserver.HTTPService, p *Platform) {
	repo := repositories.NewPaymentMethodRepository(p.DB)
	listUC := usecases.NewListPaymentMethodsUseCase(repo)
	handler := handlers.NewPaymentMethodHandler(listUC)

	group := svc.Group("api/payment-methods")
	group.GET("", handler.ListPaymentMethods)
}
