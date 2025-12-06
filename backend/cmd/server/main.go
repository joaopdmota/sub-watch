package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sub-watch/application"
	"sub-watch/application/config"
	server "sub-watch/infra/http/router"
	"sub-watch/infra/otel"

	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

func main() {
	envs := config.LoadEnvs()
	application.InitializeDependencies()

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

	httpService := server.NewHTTPService(strconv.Itoa(envs.ApiPort))

	go func() {
		if err := httpService.Start(); err != nil {
			log.Fatalf("Erro ao iniciar o servidor: %v", err)
		}
	}()

	<-ctx.Done()

	if err := httpService.Stop(context.Background()); err != nil {
		log.Fatalf("Erro ao parar o servidor: %v", err)
	}
}
