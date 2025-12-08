package otel

import (
	"context"
	"os"
	"sub-watch/application/services"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Init(ctx context.Context, logger services.Logger) func() {
	endpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")

	if endpoint == "" {
		logger.Warn("[OTEL] OTEL_EXPORTER_OTLP_ENDPOINT empty, OTEL disabled")
		return func() {}
	}

conn, err := grpc.NewClient(
		endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Error("[OTEL] failed to create gRPC client:", err)
		return func() {}
	}

	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		logger.Error("[OTEL] failed to create exporter: ", err)
		return func() {}
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(os.Getenv("OTEL_SERVICE_NAME")),
		)),
	)

	otel.SetTracerProvider(tp)

	return func() {
		logger.Info("[OTEL] Shutdown")
		_ = tp.Shutdown(ctx)
	}
}
