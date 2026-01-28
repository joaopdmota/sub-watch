package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sub-watch-backend/internal/application/config"
	infra_init "sub-watch-backend/internal/infra/init"
	"sub-watch-backend/internal/infra/logger"
	"sub-watch-backend/internal/infra/otel"

	_ "sub-watch-backend/docs"
)

// @title           Sub Watch API
// @version         1.0
// @description     This is the API server for Sub Watch application.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
	func main() {
		envs := config.LoadEnvs()
		logger := logger.New()
		shutdownOtel := func() {}

		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
		defer stop()

		if envs.OtelEnabled {
			shutdownOtel = otel.Init(ctx, logger)
		}
		defer shutdownOtel()

		httpService := infra_init.InitializeDependencies(envs, logger)

		go func() {
			if err := httpService.Start(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("Error starting the server: %v", err)
			}
		}()

		<-ctx.Done()

		if err := httpService.Stop(context.Background()); err != nil {
			log.Fatalf("Error stopping the server: %v", err)
		}
	}
