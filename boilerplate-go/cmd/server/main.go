package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sub-watch/application"
	"sub-watch/application/config"
	"sub-watch/infra/logger"
	"sub-watch/infra/otel"
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
