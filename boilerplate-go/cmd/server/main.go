package main

import (
	"boilerplate-go/internal/api"
	"boilerplate-go/internal/application"
	"boilerplate-go/internal/application/config"
	"boilerplate-go/internal/infra/logger"
	"boilerplate-go/internal/infra/otel"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

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

	httpService := application.InitializeDependencies(envs)
	api.RegisterRoutes(httpService)

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
