package main

import (
	"boilerplate-go/internal/application/config"
	"boilerplate-go/internal/infra"
	"boilerplate-go/internal/infra/grpc"
	http_infra "boilerplate-go/internal/infra/http"
	"boilerplate-go/internal/infra/logger"
	"boilerplate-go/internal/infra/otel"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
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

	httpService := infra.InitializeDependencies(envs)
	http_infra.RegisterRoutes(httpService)

	grpcService := grpc.NewGRPCService(strconv.Itoa(envs.GRPCPort))

	go func() {
		if err := httpService.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting the HTTP server: %v", err)
		}
	}()

	go func() {
		if err := grpcService.Start(); err != nil {
			log.Fatalf("Error starting the gRPC server: %v", err)
		}
	}()

	<-ctx.Done()
	fmt.Println("Shutting down servers...")

	grpcService.Stop()

	if err := httpService.Stop(context.Background()); err != nil {
		log.Fatalf("Error stopping the server: %v", err)
	}
}
