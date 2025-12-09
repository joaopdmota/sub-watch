# sub-watch-backend

Backend em Go pensado para você **subir uma API rápido** sem abrir mão de boa arquitetura, DX e observabilidade opcional.

---

## ✨ Funcionalidades

- **Arquitetura limpa**  
  - Separação clara entre `cmd`, `application`, `domain`, `infra` e `api`
  - Fácil de testar, manter e evoluir

- **Servidor HTTP desacoplado**  
  - Camada HTTP isolada em `infra/http` (router, handlers, middlewares, webserver)
  - `application` e `domain` não sabem nada de HTTP

- **Config centralizada**  
  - Leitura e validação de envs em `application/config/env.go`

- **Logger estruturado**  
  - Implementação com `slog` em `infra/logger`
  - Interface `Logger` em `application/services` para manter o domínio desacoplado

- **OpenTelemetry pronto para uso (mas opcional)**  
  - Integração em `infra/otel`
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
cd sub-watch-backend
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
SERVICE_NAME=sub-watch-backend

# Observabilidade (opcional)
OTEL_ENABLED=false
OTEL_EXPORTER_OTLP_ENDPOINT=http://otel-collector:4317
OTEL_SERVICE_NAME=sub-watch-backend
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
| `SERVICE_NAME`                | Nome lógico do serviço                             | `sub-watch-backend`                            |
| `OTEL_ENABLED`                | Liga/desliga OTEL (`true` / `false`)               | `false`                                     |
| `OTEL_EXPORTER_OTLP_ENDPOINT` | Endpoint OTLP do collector                         | `http://otel-collector:4317` (exemplo)      |
| `OTEL_SERVICE_NAME`           | Nome do serviço nos traces                         | `sub-watch-backend`                            |
| `OTEL_RESOURCE_ATTRIBUTES`    | Atributos extras de resource do OTEL               | `service.version=1.0.0,service.environment=local` |

Você pode adicionar outras envs de domínio conforme for evoluindo o projeto (DB, Redis, etc.).

---

## 📁 Estrutura de Pastas

Visão geral:

```text
sub-watch-backend/
├── .build/
│   ├── dev/
│   │   └── Dockerfile.dev      # Ambiente de desenvolvimento (Air, Go, etc.)
│   └── prod/                   # Dockerfiles de produção (a definir)
├── api/                        # DTOs, contratos de entrada/saída, schemas
├── application/
│   ├── config/
│   │   ├── env.go              # Carregamento de envs
│   │   └── env_test.go
│   ├── domain/                 # Entidades e regras de negócio puras
│   ├── services/
│   │   └── logger.go           # Interface de logger
│   └── usecases/
│       ├── dependencies.go     # Composition root / injeção de dependências
│       └── error.go            # Tipos de erro da aplicação
├── cmd/
│   └── server/
│       └── main.go             # Entrypoint da API
├── docs/                       # Documentação técnica (diagramas, notas, etc.)
├── infra/
│   ├── http/
│   │   ├── client/             # Clientes HTTP externos (se houver)
│   │   ├── handlers/           # Handlers HTTP (camada de borda)
│   │   ├── middlewares/        # Middlewares (logger, recovery, etc.)
│   │   ├── router/             # Registro de rotas
│   │   └── webserver/          # Server HTTP (start/stop, graceful shutdown)
│   ├── logger/
│   │   └── logger.go           # Implementação concreta do logger (slog)
│   └── otel/
│       └── otel.go             # Integração com OpenTelemetry
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

Um fluxo típico de requisição HTTP na sua API pode ser:

```text
1. [HTTP Request] → Handler em infra/http/handlers
2. Handler:
   - valida/parsa entrada
   - converte para DTO de usecase
3. Handler chama → UseCase em application/usecases
4. UseCase:
   - aplica regra de negócio
   - chama interfaces de serviços/repos (definidas em application/services)
5. Implementações concretas em infra/* executam:
   - chamadas HTTP externas
   - acesso a banco de dados
   - etc.
6. UseCase retorna DTO de saída
7. Handler converte para JSON → responde para o cliente
```

O domínio (`application/domain`) não conhece HTTP, banco, nem nada de infra.

---

## 🧾 Logger

Interface de logger em `application/services/logger.go`:

```go
type Logger interface {
    Info(msg string, kv ...any)
    Warn(msg string, kv ...any)
    Error(msg string, kv ...any)
    Debug(msg string, kv ...any)
}
```

Implementação concreta com `slog` em `infra/logger/logger.go`:

```go
type SlogLogger struct {
    l *slog.Logger
}

func New() *SlogLogger {
    return &SlogLogger{
        l: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
    }
}

func (s *SlogLogger) Info(msg string, kv ...any)  { s.l.Info(msg, kv...) }
func (s *SlogLogger) Warn(msg string, kv ...any)  { s.l.Warn(msg, kv...) }
func (s *SlogLogger) Error(msg string, kv ...any) { s.l.Error(msg, kv...) }
func (s *SlogLogger) Debug(msg string, kv ...any) { s.l.Debug(msg, kv...) }
```

Isso permite:

- **domínio e usecases** dependerem apenas da interface `Logger`;
- trocar a implementação (slog → zap → zerolog) sem alterar regra de negócio.

---

## 📡 Observabilidade (OpenTelemetry)

A integração com OTEL está em `infra/otel/otel.go`.

Pontos chave:

- Controlada por `OTEL_ENABLED`:
  - `false` → não tenta conectar, só loga que está desabilitado
  - `true` → tenta inicializar tracing
- Se não conseguir conectar no collector:
  - loga o erro
  - **não impede a aplicação de subir**  
    (observabilidade é “best effort”, não requisito de vida ou morte)
- Exemplo de uso no `main.go`:

```go
ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
defer stop()

shutdownOtel := func() {}

if envs.OtelEnabled {
    shutdownOtel = otel.Init(ctx)
}

defer shutdownOtel()
```

Depois é só instrumentar handlers/usecases com spans, se quiser.

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