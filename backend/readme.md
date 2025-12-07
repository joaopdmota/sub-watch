# Sub-Watch Backend

Este repositório contém a implementação do backend do serviço `sub-watch`. O projeto fornece uma estrutura robusta para o desenvolvimento de serviços em Go, incluindo configurações pré-definidas para observabilidade, containerização e desenvolvimento local.

## Funcionalidades Principais

- **Servidor Web**: Construído sobre o framework [Echo](https://echo.labstack.com/), oferecendo alta performance e facilidade de extensão.
- **Observabilidade (OpenTelemetry)**: Integração nativa para coleta de métricas e traces, garantindo visibilidade sobre o comportamento da aplicação.
- **Tracing Distribuído**: Configuração pronta para envio de traces ao Zipkin/Jaeger via OpenTelemetry Collector.
- **Ambiente de Desenvolvimento**: 
  - Docker Compose para orquestração de serviços locais.
  - Suporte a Hot Reload utilizando [Air](https://github.com/cosmtrek/air) (configurado no Dockerfile de dev).

---

## Estrutura de Serviços (Docker Compose)

O ambiente local é composto pelos seguintes serviços definidos no `compose.yaml`:

### 1. **otel-collector**
O OpenTelemetry Collector é responsável por receber, processar e exportar dados de telemetria (métricas e traces) da aplicação.

- **Imagem**: `otel/opentelemetry-collector-contrib`
- **Portas e Protocolos**:
  - `4317`: OTLP gRPC (Recebimento de dados da aplicação).
  - `4318`: OTLP HTTP.
  - `1888`: Métricas internas do collector.
  - `8888`: Porta de métricas para Prometheus (se configurado).
  - `13133`: Health Check.
  - `55679`: Z-Pages (Debug legacy).
- **Configuração**: Mapeia o arquivo `./otel-collector-config.yaml` para `/etc/otelcol-contrib/config.yaml`.
- **Dependências**: Aguarda o início do serviço `zipkin`.

### 2. **zipkin**
Serviço de armazenamento e visualização de traces distribuídos. Útil para depurar latência e fluxo de requisições.

- **Imagem**: `openzipkin/zipkin`
- **Interface Web**: Acessível em `http://localhost:9411`.
- **Armazenamento**: Configurado para `MEM` (em memória), ou seja, os dados são perdidos ao reiniciar o container.

### 3. **server**
O serviço principal da aplicação backend.

- **Porta**: `8080` (Acessível via `http://localhost:8080`).
- **Desenvolvimento**:
  - Mapeia o diretório atual (`.`) para dentro do container, permitindo edição de código em tempo real.
  - Utiliza `air` para recompilar e reiniciar a aplicação automaticamente ao detectar mudanças nos arquivos.
- **Variáveis de Ambiente**: Carregadas a partir do arquivo `.env`.

---

## Como Utilizar

### Pré-requisitos
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go 1.23+](https://go.dev/dl/) (Opcional, caso queira rodar fora do Docker)

### Passo a Passo

1. **Configuração de Ambiente**:
   Crie um arquivo `.env` na raiz do projeto com base no `.env.example` (se houver) ou definindo as variáveis necessárias (ex: `API_PORT=8080`).

2. **Configuração do Collector**:
   Garanta que o arquivo `otel-collector-config.yaml` esteja presente na raiz e configurado corretamente.

3. **Iniciando os Serviços**:
   Execute o comando abaixo para construir as imagens e subir o ambiente:
   ```bash
   docker-compose up --build
   ```

4. **Acessando a Aplicação**:
   - API: `http://localhost:8080`
   - Zipkin (Traces): `http://localhost:9411`

## Arquitetura do Backend

O projeto segue os princípios de **Clean Architecture** (Arquitetura Limpa), garantindo separação de responsabilidades, testabilidade e independência de frameworks externos.

### Estrutura de Camadas

```
backend/
├── cmd/                    # Camada de Entrada
│   └── server/            # Ponto de inicialização da aplicação
├── application/           # Camada de Aplicação (Regras de Negócio)
│   ├── domain/           # Entidades de Domínio (User, Subscription, etc.)
│   ├── services/         # Serviços de Aplicação (orquestração)
│   ├── usecases/         # Casos de Uso (ações específicas)
│   ├── config/           # Configurações da aplicação
│   └── dependencies.go   # Injeção de Dependências
├── infra/                # Camada de Infraestrutura (Detalhes Técnicos)
│   ├── database/         # Adaptadores de Banco de Dados
│   ├── http/             # Servidor HTTP, Handlers, Middlewares
│   ├── repositories/     # Implementações de Repositórios
│   └── otel/             # Observabilidade (OpenTelemetry)
├── docs/                 # Documentação Swagger/OpenAPI (auto-gerada)
└── api/                  # Exemplos de requisições HTTP (.http files)
```

### Camadas e Responsabilidades

#### 1. **Domain (Domínio)**
- **Localização**: `application/domain/`
- **Responsabilidade**: Define as entidades principais do sistema (ex: `User`, `Subscription`, `PaymentMethod`)
- **Características**:
  - Sem dependências externas
  - Modelos de dados puros (structs Go)
  - Representa conceitos de negócio

#### 2. **Use Cases (Casos de Uso)**
- **Localização**: `application/usecases/`
- **Responsabilidade**: Implementa regras de negócio específicas (ex: `ListUsersUseCase`, `GetUserUseCase`)
- **Características**:
  - Orquestra fluxos de dados
  - Depende de `Services` (não de detalhes técnicos)
  - Retorna DTOs (`UserOutput`) ao invés de entidades diretas

#### 3. **Services (Serviços de Aplicação)**
- **Localização**: `application/services/`
- **Responsabilidade**: Interface entre Use Cases e Repositories
- **Características**:
  - Define contratos (interfaces) para operações de dados
  - Ex: `UserService.GetAllUsers()`, `UserService.GetUserByID()`
  - Permite trocar implementações facilmente

#### 4. **Repositories (Repositórios)**
- **Localização**: `infra/repositories/`
- **Responsabilidade**: Acesso a dados (comunicação com banco)
- **Características**:
  - Implementa interfaces definidas em `Services`
  - Usa abstração `database.Database` (não SQL direto)
  - Ex: `UserRepository.FindAll()`, `UserRepository.FindByID()`

#### 5. **Database Adapters (Adaptadores de Banco)**
- **Localização**: `infra/database/`
- **Responsabilidade**: Abstrai operações de banco específicas (PostgreSQL)
- **Características**:
  - Interface `Database` com métodos genéricos (`FindAll`, `FindByID`)
  - `PostgresAdapter` implementa essa interface
  - Gera queries SQL internamente (repositórios não sabem que é SQL)

#### 6. **HTTP Handlers**
- **Localização**: `infra/http/handlers/`
- **Responsabilidade**: Recebe requisições HTTP e delega para Use Cases
- **Características**:
  - Valida entrada
  - Converte dados HTTP (JSON) para objetos Go
  - Retorna respostas (status codes, erros)

### Fluxo de uma Requisição (Exemplo: GET /users/123)

```
1. [HTTP Request] → Handler (GetUser)
2. Handler extrai ID do parâmetro
3. Handler chama → GetUserUseCase.Execute(ctx, id)
4. UseCase chama → UserService.GetUserByID(ctx, id)
5. Service delega → UserRepository.FindByID(ctx, id)
6. Repository usa → Database.FindByID(ctx, "users", id)
7. PostgresAdapter executa → SELECT * FROM users WHERE id = $1
8. Dados retornam na ordem inversa até o Handler
9. Handler retorna JSON → Cliente HTTP
```

### Princípios Aplicados

- **Inversão de Dependências**: Camadas internas não conhecem as externas
- **Abstração de Banco**: Repositories usam interface `Database`, não SQL direto
- **Testabilidade**: Cada camada pode ser testada isoladamente com mocks
- **Desacoplamento**: Trocar PostgreSQL por MongoDB afeta apenas `infra/database`

### Injeção de Dependências

O arquivo `application/dependencies.go` inicializa todas as dependências na ordem correta:

```go
Database → Repository → Service → UseCase → Handler → Router
```

## Documentação da API (Swagger)

A API possui documentação interativa via Swagger/OpenAPI, acessível em:
- **Swagger UI**: `http://localhost:8080/swagger/index.html`

### Atualizando a Documentação

Após adicionar ou modificar endpoints com anotações Swagger, regenere a documentação executando:

```bash
make docs
```

Este comando:
1. Gera os arquivos de documentação em `docs/` com base nas anotações do código
2. Remove campos incompatíveis automaticamente
3. Permite que a interface Swagger reflita as mudanças mais recentes

**Nota**: Certifique-se de que o `swag` CLI está instalado (`go install github.com/swaggo/swag/cmd/swag@latest`)
