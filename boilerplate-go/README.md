# boilerplate-go

Backend em Go pensado para você **subir uma API rápido** sem abrir mão de boa arquitetura, DX e observabilidade opcional.

---

## ✨ Funcionalidades

- **Arquitetura limpa**  
  - Separação clara entre `cmd`, `application`, `domain`, `infra` e `pkg`
  - Fácil de testar, manter e evoluir

- **Servidor HTTP desacoplado**  
  - Camada HTTP isolada em `infra/http` (router, handlers, middlewares, webserver)
  - `application` e `domain` não sabem nada de HTTP

- **Config centralizada**  
  - Leitura e validação de envs em `application/config/env.go`

- **Logger estruturado**  
  - Interface de logger no domínio (ex.: via `pkg/logger`)
  - Implementação concreta em `infra/logger` (quando aplicável)

- **Providers reutilizáveis em `pkg/`**  
  - `pkg/id`: geração de IDs (UUID)
  - `pkg/hash`: hashing seguro de senha (bcrypt)
  - `pkg/date`: provider de datas testável (`Now()` injetável)
  - `pkg/logger`: abstração de logger reutilizável entre serviços

- **OpenTelemetry pronto para uso (mas opcional)**  
  - Integração em `infra/otel` (quando configurado)
  - Controle via `OTEL_ENABLED`
  - Se o collector estiver fora do ar, a app **continua funcionando**

- **Ambiente de desenvolvimento com Docker + Air**  
  - Hot reload dentro do container
  - `compose.yaml` e `.build/dev/Dockerfile.dev` já configurados

- **Makefile para rotina diária**  
  - Subir/parar serviço, ver logs, rodar testes, `go mod tidy`, etc.

---

## 🚀 Início Rápido

### Pré-requisitos

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- (Opcional) [Go 1.21+](https://go.dev/dl/) se quiser rodar fora do Docker

---

### 1. Clonar o repositório

```bash
git clone <url-do-repo>
cd boilerplate-go
```

---

### 2. Configurar variáveis de ambiente

Crie seu `.env` a partir do exemplo:

```bash
cp .env.example .env
```

Edite o `.env` com os valores desejados. Exemplo mínimo:

```env
API_PORT=8080
SERVICE_NAME=boilerplate-go

# Observabilidade (opcional)
OTEL_ENABLED=false
OTEL_EXPORTER_OTLP_ENDPOINT=http://otel-collector:4317
OTEL_SERVICE_NAME=boilerplate-go
OTEL_RESOURCE_ATTRIBUTES=service.version=1.0.0,service.environment=local
```

> Em desenvolvimento é comum deixar `OTEL_ENABLED=false`.  
> Em produção você liga e aponta para seu collector.

---

### 3. Subir o servidor com Docker + Air

```bash
docker compose up --build
# ou, se configurado:
# make up
```

A aplicação sobe com hot reload dentro do container.

Acesse:

- API (exemplo de root):  
  `http://localhost:8080`

*(Você pluga aqui os endpoints da sua aplicação.)*

---

### 4. Ver logs do servidor

```bash
docker compose logs -f server
# ou
# make server-logs
```

---

### 5. Rodar localmente sem Docker (opcional)

```bash
go mod download
go run cmd/server/main.go
```

---

## ⚙️ Variáveis de ambiente

Carregadas em `application/config/env.go`.

| Variável                      | Descrição                                           | Default (sugerido)                          |
| ----------------------------- | --------------------------------------------------- | ------------------------------------------- |
| `API_PORT`                    | Porta em que o servidor HTTP escuta                | `8080`                                      |
| `SERVICE_NAME`                | Nome lógico do serviço                             | `boilerplate-go`                            |
| `OTEL_ENABLED`                | Liga/desliga OTEL (`true` / `false`)               | `false`                                     |
| `OTEL_EXPORTER_OTLP_ENDPOINT` | Endpoint OTLP do collector                         | `http://otel-collector:4317` (exemplo)      |
| `OTEL_SERVICE_NAME`           | Nome do serviço nos traces                         | `boilerplate-go`                            |
| `OTEL_RESOURCE_ATTRIBUTES`    | Atributos extras de resource do OTEL               | `service.version=1.0.0,service.environment=local` |

Você pode adicionar outras envs de domínio conforme for evoluindo o projeto (DB, Redis, etc.).

---

## 📁 Estrutura de Pastas

Visão geral (adaptar para sua estrutura real de microserviço):

```text
boilerplate-go/
├── .build/
│   ├── dev/
│   │   └── Dockerfile.dev      # Ambiente de desenvolvimento (Air, Go, etc.)
│   └── prod/                   # Dockerfiles de produção (a definir)
├── application/
│   ├── config/
│   │   ├── env.go              # Carregamento de envs
│   │   └── env_test.go
│   ├── domain/                 # Entidades e regras de negócio puras
│   └── usecases/               # Casos de uso da aplicação
│   └── logger.go               # Abstrações de logger
├── cmd/
│   └── server/
│       └── main.go             # Entrypoint da API
├── infra/
│   ├── http/
│   │   ├── client/             # Clientes HTTP externos (se houver)
│   │   ├── handlers/           # Handlers HTTP (camada de borda)
│   │   ├── middlewares/        # Middlewares (logger, recovery, etc.)
│   │   └── webserver/          # Server HTTP (start/stop, graceful shutdown)
│   ├── database/               # Interfaces e adapters de banco (ex.: PostgresAdapter)
│   ├── logger/                 # Implementação concreta do logger
│   └── otel/                   # Integração com OpenTelemetry
├── pkg/
│   ├── date/                   # Provider de datas (ex.: Now())
│   ├── hash/                   # Hash de senha (bcrypt, etc.)
│   ├── id/                     # Gerador de IDs (UUID)
├── tmp/                        # Artefatos temporários (binário gerado pelo Air)
├── .air.toml                   # Configuração do Air (hot reload)
├── .env                        # Env local (não versionar)
├── .env.example                # Modelo de env
├── compose.yaml                # Docker Compose para dev
├── go.mod
├── go.sum
├── Makefile
└── readme.md
```

---

## 🔌 Fluxo de uma requisição (visão conceitual)

```text
1. [HTTP Request] → Handler em infra/http/handlers
2. Handler:
   - valida/parsa entrada
   - converte para DTO de usecase
3. Handler chama → UseCase em application/usecases
4. UseCase:
   - aplica regra de negócio
   - chama interfaces de serviços/repos
5. Implementações concretas em infra/* executam:
   - chamadas HTTP externas
   - acesso a banco de dados
   - logging, tracing, etc.
6. UseCase retorna DTO de saída
7. Handler converte para JSON → responde para o cliente
```

O domínio (`application/domain`) não conhece HTTP, banco, nem nada de infra.

---

## 🧾 Logger

Interface de logger no domínio (exemplo):

```go
type Logger interface {
    Info(msg string, kv ...any)
    Warn(msg string, kv ...any)
    Error(msg string, kv ...any)
    Debug(msg string, kv ...any)
}
```

Implementações concretas podem viver em `infra/logger` e/ou `pkg/logger`, usando `slog`, `zap` etc., mantendo o domínio desacoplado.

---

## 📡 Observabilidade (OpenTelemetry)

Quando configurado, a integração com OTEL fica em `infra/otel`.

Pontos chave:

- Controlada por `OTEL_ENABLED`
- Se não conseguir conectar no collector:
  - loga o erro
  - **não impede a aplicação de subir**

---

## 🧰 Comandos úteis (Makefile)

```bash
make up            # Sobe server com Docker Compose
make down          # Derruba containers
make server-logs   # Tail nos logs do servidor
make tidy          # go mod tidy dentro do container
make test          # go test ./...
```

---

Este boilerplate foi pensado para servir de base para microserviços Go (como o SubWatch) com foco em **claridade de arquitetura**, **testabilidade** e **reutilização** de utilitários em `pkg/`.
