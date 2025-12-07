package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sub-watch/application"
	"sub-watch/application/config"
	"sub-watch/infra/otel"

	_ "sub-watch/docs"
)

// @title Sub-Watch API
// @version 1.0
// @description Backend API for Sub-Watch application.
// @host localhost:8080
// @BasePath /
func main() {
	envs := config.LoadEnvs()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	shutdown, err := otel.SetupOTelSDK(ctx, envs.ServiceName)
	if err != nil {
		log.Fatalf("Erro ao iniciar observabilidade: %v", err)
	}
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Printf("Erro ao desligar OTel SDK: %v", err)
		}
	}()

	httpService := application.InitializeDependencies(envs)

	go func() {
		if err := httpService.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erro ao iniciar o servidor: %v", err)
		}
	}()

	<-ctx.Done()

	if err := httpService.Stop(context.Background()); err != nil {
		log.Fatalf("Erro ao parar o servidor: %v", err)
	}
}
