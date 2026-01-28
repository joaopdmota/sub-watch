# boilerplate-go

Go Backend designed for **rapid API deployment** without compromising on architecture, DX, and optional observability.

---

## ✨ Features

- **Clean Architecture**
  - Clear separation between `cmd`, `internal/application`, `internal/domain`, `internal/infra`, and `internal/pkg`.
  - Easy to test, maintain, and evolve.

- **Dual Protocol Support (REST & gRPC)**
  - **REST**: Powered by Echo, isolated in `internal/infra/http`.
  - **gRPC**: Built-in support with reflection, isolated in `internal/infra/grpc`.
  - Unified orchestration allows running both simultaneously.

- **Event-Driven Ready (Kafka)**
  - Dedicated **stream processor** entrypoint in `cmd/stream`.
  - Resilient Kafka subscriber implemented in `internal/infra/kafka` using `segmentio/kafka-go`.

- **Centralized Configuration**
  - Environment variable loading and validation in `internal/application/config/env.go`.

- **Structured Logging**
  - Dedicated implementation in `internal/infra/logger`.

- **Reusable Providers in `internal/pkg/`**
  - `internal/pkg/id`: UUID generation.
  - `internal/pkg/hash`: Secure password hashing (bcrypt).
  - `internal/pkg/date`: Testable date provider (injectable `Now()`).

- **Ready-to-use Observability (OpenTelemetry)**
  - Integration in `internal/infra/otel`.
  - Controlled via `OTEL_ENABLED`.
  - Application remains functional even if the collector is down.

- **Development Environment with Docker + Air**
  - Hot reload inside the container.
  - Pre-configured `compose.yaml` and `.build/dev/Dockerfile.dev`.

- **Unified Tooling with Makefile**
  - Commands for building, testing, proto generation, and more.

---

## 🚀 Quick Start

### Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- (Optional) [Go 1.24+](https://go.dev/dl/)

---

### 1. Clone the repository

```bash
git clone <repo-url>
cd boilerplate-go
```

---

### 2. Configure Environment Variables

Create your `.env` from the example:

```bash
cp .env.example .env
```

Key configuration variables:

```env
API_PORT=8080
GRPC_PORT=9090
SERVICE_NAME=boilerplate-go

# Kafka Configuration
KAFKA_BROKERS=localhost:9092
KAFKA_TOPIC=my-topic
KAFKA_GROUP_ID=my-group

# Observability (optional)
OTEL_ENABLED=false
```

---

### 3. Generate gRPC Code

```bash
make proto
```
*Note: If you don't have protoc installed locally, the CI or a development container should handle this.*

---

### 4. Run the Services

#### Run the API (HTTP & gRPC)
```bash
make up
# or
docker compose up --build server
```

#### Run the Stream Processor (Kafka Subscriber)
```bash
docker compose up --build stream
```

---

## ⚙️ Environment Variables

Defined in `internal/application/config/env.go`.

| Variable | Description | Default/Example |
| :--- | :--- | :--- |
| `API_PORT` | HTTP server port | `8080` |
| `GRPC_PORT` | gRPC server port | `9090` |
| `SERVICE_NAME` | Logical service name | `boilerplate-go` |
| `KAFKA_BROKERS` | Kafka broker list (comma separated) | `localhost:9092` |
| `KAFKA_TOPIC` | Topic to subscribe to | `subwatch-messages` |
| `KAFKA_GROUP_ID` | Consumer group ID | `subwatch-group` |
| `OTEL_ENABLED` | Enable/Disable OTEL (`true`/`false`) | `false` |

---

## 📁 Folder Structure

```text
boilerplate-go/
├── api/
│   ├── proto/               # Protobuf definitions
│   └── test.http            # REST Specs & snippets
├── cmd/
│   ├── server/
│   │   └── main.go          # Main entrypoint (HTTP + gRPC)
│   └── stream/
│       └── main.go          # Stream entrypoint (Kafka subscriber)
├── internal/
│   ├── api/                 # API routing and registration
│   ├── application/         # Use Cases & Business Logic
│   │   └── config/          # App configuration
│   ├── domain/              # Entities & Repository Interfaces
│   ├── infra/               # Infrastructure Adapters
│   │   ├── grpc/            # gRPC server & generated code
│   │   ├── http/            # REST webserver, handlers, middlewares
│   │   ├── kafka/           # Kafka subscriber implementation
│   │   ├── logger/          # Structured logging implementation
│   │   └── otel/            # OpenTelemetry integration
│   └── pkg/                 # Shared utilities (date, hash, id)
├── Makefile                 # Automation tasks
└── compose.yaml             # Local development orchestration
```

---

## 🧰 Useful Commands

```bash
make up            # Start servers with Docker Compose
make down          # Stop containers
make proto         # Generate gRPC code from .proto files
make test          # Run all tests
make tidy          # Run go mod tidy
```

---

This boilerplate is designed to be the foundation for scalable Go microservices, focusing on **architectural clarity**, **testability**, and **DX**.
