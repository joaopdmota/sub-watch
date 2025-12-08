# Observability Stack

This project sets up a simple observability stack using Docker Compose. It includes an OpenTelemetry Collector and a Jaeger instance for collecting and visualizing traces.

## Services

- **`otel-collector`**: Receives telemetry data (traces, metrics, logs) via OTLP (OpenTelemetry Protocol) over gRPC and HTTP. It then processes and exports this data to Jaeger.
- **`jaeger`**: A distributed tracing system. It receives trace data from the OTEL Collector and provides a web UI to visualize them.

## Getting Started

To start the observability stack, run the following command:

```bash
docker compose up -d
```

This will start both the `otel-collector` and `jaeger` services in detached mode.

## Endpoints

Once the services are running, you can access the following endpoints:

- **Jaeger UI**: [http://localhost:16686](http://localhost:16686)
- **OTLP gRPC Receiver**: `localhost:4317`
- **OTLP HTTP Receiver**: `localhost:4318`

## Configuration

The OpenTelemetry Collector is configured via the `otel-collector-config.yaml` file. By default, it's set up to:

- **Receive** data using the OTLP protocol on ports `4317` (gRPC) and `4318` (HTTP).
- **Process** the data in batches.
- **Export** the trace data to the Jaeger instance at `jaeger:4317`.
