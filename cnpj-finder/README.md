# Consulta CPF/CNPJ API

API simples e inteligente para consulta de documentos brasileiros (CPF e CNPJ) na Receita Federal através da [ReceitaWS](https://receitaws.com.br/).

## ✨ Funcionalidades

- **Consulta Unificada**: Um único endpoint que detecta automaticamente se o documento é CPF ou CNPJ
- **Validação Automática**: Valida formato baseado no número de dígitos (11 para CPF, 14 para CNPJ)
- **Documentação Swagger**: Interface interativa para testar a API
- **Observabilidade**: Integração com OpenTelemetry para traces e métricas

---

## 🚀 Início Rápido

### Pré-requisitos
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go 1.23+](https://go.dev/dl/) (Opcional, para desenvolvimento local)

### Como Usar

1. **Clone o repositório** e navegue até a pasta do projeto

2. **Configure as variáveis de ambiente**:
   ```bash
   cp .env.example .env
   ```
   
   Edite o `.env` com as configurações desejadas:
   ```env
   API_PORT=8080
   SERVICE_NAME=doc-searcher
   ```

3. **Inicie os serviços**:
   ```bash
   make run
   # ou
   docker-compose up --build
   ```

4. **Acesse a API**:
   - **API**: http://localhost:8080
   - **Swagger UI**: http://localhost:8080/swagger/index.html
   - **Health Check**: http://localhost:8080/health
   - **Zipkin (Traces)**: http://localhost:9411

---

## 📡 API

### `GET /consulta/:numero`

Consulta dados de um CPF ou CNPJ na Receita Federal. O tipo de documento é **detectado automaticamente** baseado no número de dígitos.

**Parâmetros:**
- `numero` (path): Número do CPF ou CNPJ, com ou sem formatação
  - **CPF**: 11 dígitos (ex: `12345678900` ou `123.456.789-00`)
  - **CNPJ**: 14 dígitos (ex: `06990590000123` ou `06.990.590/0001-23`)

---

### Exemplo 1: Consultar CNPJ

**Requisição:**
```bash
curl http://localhost:8080/consulta/06990590000123
```

**Resposta (200 OK):**
```json
{
  "tipo": "CNPJ",
  "cnpj": "06.990.590/0001-23",
  "razao_social": "GOOGLE BRASIL INTERNET LTDA",
  "nome_fantasia": "Google",
  "situacao_cadastral": "2",
  "descricao_situacao_cadastral": "Ativa",
  "data_inicio_atividade": "03/11/2005",
  "cnae_fiscal": "6319400",
  "cnae_fiscal_descricao": "Portais, provedores de conteúdo e outros serviços de informação na internet",
  "logradouro": "AV BRIGADEIRO FARIA LIMA",
  "numero": "3477",
  "bairro": "ITAIM BIBI",
  "cep": "04538133",
  "uf": "SP",
  "municipio": "SAO PAULO",
  "telefone": "1121395000",
  "capital_social": "200000000.00",
  "porte": "05",
  "porte_descricao": "Demais"
}
```

---

### Exemplo 2: Consultar CPF

**Requisição:**
```bash
curl http://localhost:8080/consulta/12345678900
```

**Resposta (200 OK):**
```json
{
  "tipo": "CPF",
  "cpf": "123.456.789-00",
  "nome": "João da Silva",
  "situacao": "Regular",
  "data_nascimento": "01/01/1990",
  "sexo": "M"
}
```

---

### Respostas de Erro

- **400 Bad Request**: Documento com formato inválido
  ```json
  {
    "message": "número inválido: deve conter 11 dígitos (CPF) ou 14 dígitos (CNPJ)"
  }
  ```

- **404 Not Found**: Documento não encontrado na base da Receita

- **500 Internal Server Error**: Erro ao consultar a API externa

---

## 🏗 Arquitetura

O projeto segue os princípios de **Clean Architecture**, mantendo separação de responsabilidades:

```
.
├── cmd/server/          # Ponto de entrada da aplicação
├── application/         # Camada de Aplicação
│   ├── config/         # Configurações e variáveis de ambiente
│   ├── domain/         # Entidades de domínio (CNPJ, CPF)
│   ├── services/       # Interfaces de serviços
│   └── usecases/       # Casos de uso (GetCNPJUseCase, GetCPFUseCase)
└── infra/              # Camada de Infraestrutura
    ├── http/
    │   ├── clients/    # Cliente HTTP para ReceitaWS
    │   ├── handlers/   # Handler unificado de documentos
    │   ├── middlewares/# Middlewares (logger, etc.)
    │   └── router/     # Configuração de rotas
    └── otel/           # Observabilidade (OpenTelemetry)
```

### Fluxo de uma Requisição

```
1. [HTTP Request] → DocumentHandler.Get()
2. Handler detecta tipo (11 dígitos = CPF, 14 dígitos = CNPJ)
3. Handler chama use case apropriado:
   - GetCPFUseCase.Execute() para CPF
   - GetCNPJUseCase.Execute() para CNPJ
4. UseCase valida formato
5. UseCase chama ReceitaClient.GetByCPF() ou GetByCNPJ()
6. ReceitaClient faz requisição HTTP para ReceitaWS API
7. Dados retornam com campo "tipo" indicando CPF ou CNPJ
8. Handler retorna JSON unificado → Cliente HTTP
```

---

## 💻 Desenvolvimento

### Comandos Disponíveis

```bash
make run          # Inicia a aplicação com Docker Compose
make build        # Rebuild das imagens Docker
make down         # Para todos os containers
make test         # Executa testes
make test-cover   # Executa testes com cobertura HTML
make docs         # Regenera documentação Swagger
make server-logs  # Visualiza logs do servidor
```

### Executar Localmente (sem Docker)

```bash
# Instalar dependências
go mod download

# Executar aplicação
go run cmd/server/main.go

# Executar testes
go test ./...
```

### Atualizar Documentação Swagger

Após modificar as anotações Swagger nos handlers, execute:

```bash
make docs
```

---

## 🐳 Estrutura do Docker Compose

### Serviços Disponíveis

#### **server**
- Aplicação principal
- Porta: `8080`
- Hot reload habilitado via Air

#### **otel-collector**
- Coleta métricas e traces OpenTelemetry
- Exporta para Zipkin

#### **zipkin**
- Interface para visualização de traces distribuídos
- Porta: `9411`

---

## 📊 Observabilidade

O projeto inclui integração com OpenTelemetry para:
- **Traces**: Rastreamento de requisições HTTP
- **Métricas**: Estatísticas de performance

Acesse o Zipkin em http://localhost:9411 para visualizar traces das requisições.

---

## 🛠 Tecnologias Utilizadas

- **Go 1.23**: Linguagem de programação
- **Echo**: Framework web
- **ReceitaWS API**: Fonte de dados de CPF e CNPJ
- **OpenTelemetry**: Observabilidade
- **Zipkin**: Visualização de traces
- **Swagger**: Documentação da API
- **Docker**: Containerização

---

## ⚠️ Limitações

- A API ReceitaWS é pública e possui **limite de requisições** (consulte: https://receitaws.com.br/)
- Não há autenticação nesta versão
- Dados dependem da disponibilidade da API externa
- Consultas de CPF podem ter restrições de privacidade na API pública

---

## 📝 Licença

Este projeto é open source e está disponível sob a licença MIT.


API simples para consulta de dados de CNPJ na Receita Federal através da [ReceitaWS](https://receitaws.com.br/).

## Funcionalidades

- **Consulta de CNPJ**: Endpoint único para consultar informações de empresas brasileiras
- **Validação**: Validação automática do formato do CNPJ (14 dígitos)
- **Documentação Swagger**: Interface interativa para testar a API
- **Observabilidade**: Integração com OpenTelemetry para traces e métricas

---

## Início Rápido

### Pré-requisitos
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go 1.23+](https://go.dev/dl/) (Opcional, para desenvolvimento local)

### Como Usar

1. **Clone o repositório** e navegue até a pasta do projeto

2. **Configure as variáveis de ambiente**:
   ```bash
   cp .env.example .env
   ```
   
   Edite o `.env` com as configurações desejadas:
   ```env
   API_PORT=8080
   SERVICE_NAME=cnpj-searcher
   ```

3. **Inicie os serviços**:
   ```bash
   make run
   # ou
   docker-compose up --build
   ```

4. **Acesse a API**:
   - **API**: http://localhost:8080
   - **Swagger UI**: http://localhost:8080/swagger/index.html
   - **Health Check**: http://localhost:8080/health
   - **Zipkin (Traces)**: http://localhost:9411

---

## Endpoints

### `GET /cnpj/:numero`

Consulta dados de um CNPJ na Receita Federal.

**Parâmetros:**
- `numero` (path): Número do CNPJ com ou sem formatação (ex: `06990590000123` ou `06.990.590/0001-23`)

**Exemplo de Requisição:**
```bash
curl http://localhost:8080/cnpj/06990590000123
```

**Exemplo de Resposta (200 OK):**
```json
{
  "cnpj": "06.990.590/0001-23",
  "razao_social": "GOOGLE BRASIL INTERNET LTDA",
  "nome_fantasia": "Google",
  "situacao_cadastral": "2",
  "descricao_situacao_cadastral": "Ativa",
  "data_inicio_atividade": "03/11/2005",
  "cnae_fiscal": "6319400",
  "cnae_fiscal_descricao": "Portais, provedores de conteúdo e outros serviços de informação na internet",
  "logradouro": "AV BRIGADEIRO FARIA LIMA",
  "numero": "3477",
  "bairro": "ITAIM BIBI",
  "cep": "04538133",
  "uf": "SP",
  "municipio": "SAO PAULO",
  "telefone": "1121395000",
  "capital_social": "200000000.00",
  "porte": "05",
  "porte_descricao": "Demais"
}
```

**Respostas de Erro:**
- `400 Bad Request`: CNPJ com formato inválido
- `404 Not Found`: CNPJ não encontrado
- `500 Internal Server Error`: Erro ao consultar a API externa

---

## Arquitetura

O projeto segue os princípios de **Clean Architecture**, mantendo separação de responsabilidades:

```
.
├── cmd/server/          # Ponto de entrada da aplicação
├── application/         # Camada de Aplicação
│   ├── config/         # Configurações e variáveis de ambiente
│   ├── domain/         # Entidades de domínio (CNPJ)
│   ├── services/       # Interfaces de serviços
│   └── usecases/       # Casos de uso (GetCNPJUseCase)
└── infra/              # Camada de Infraestrutura
    ├── http/
    │   ├── clients/    # Cliente HTTP para ReceitaWS
    │   ├── handlers/   # Handlers HTTP
    │   ├── middlewares/# Middlewares (logger, etc.)
    │   └── router/     # Configuração de rotas
    └── otel/           # Observabilidade (OpenTelemetry)
```

### Fluxo de uma Requisição

```
1. [HTTP Request] → CNPJHandler
2. Handler extrai número do parâmetro
3. Handler chama → GetCNPJUseCase.Execute(cnpj)
4. UseCase valida formato do CNPJ
5. UseCase chama → ReceitaClient.GetByCNPJ(cnpj)
6. ReceitaClient faz requisição HTTP para ReceitaWS API
7. Dados retornam na ordem inversa até o Handler
8. Handler retorna JSON → Cliente HTTP
```

---

## Desenvolvimento

### Comandos Disponíveis

```bash
make run          # Inicia a aplicação com Docker Compose
make build        # Rebuild das imagens Docker
make down         # Para todos os containers
make test         # Executa testes
make test-cover   # Executa testes com cobertura HTML
make docs         # Regenera documentação Swagger
make server-logs  # Visualiza logs do servidor
```

### Executar Localmente (sem Docker)

```bash
# Instalar dependências
go mod download

# Executar aplicação
go run cmd/server/main.go

# Executar testes
go test ./...
```

### Atualizar Documentação Swagger

Após modificar as anotações Swagger nos handlers, execute:

```bash
make docs
```

---

## Estrutura do Docker Compose

### Serviços Disponíveis

#### **server**
- Aplicação principal
- Porta: `8080`
- Hot reload habilitado via Air

#### **otel-collector**
- Coleta métricas e traces OpenTelemetry
- Exporta para Zipkin

#### **zipkin**
- Interface para visualização de traces distribuídos
- Porta: `9411`

---

## Observabilidade

O projeto inclui integração com OpenTelemetry para:
- **Traces**: Rastreamento de requisições HTTP
- **Métricas**: Estatísticas de performance

Acesse o Zipkin em http://localhost:9411 para visualizar traces das requisições.

---

## Tecnologias Utilizadas

- **Go 1.23**: Linguagem de programação
- **Echo**: Framework web
- **ReceitaWS API**: Fonte de dados de CNPJ
- **OpenTelemetry**: Observabilidade
- **Zipkin**: Visualização de traces
- **Swagger**: Documentação da API
- **Docker**: Containerização

---

## Limitações

- A API ReceitaWS é pública e possui **limite de requisições** (consulte: https://receitaws.com.br/)
- Não há autenticação nesta versão
- Dados dependem da disponibilidade da API externa

---

## Licença

Este projeto é open source e está disponível sob a licença MIT.


Este repositório contém a implementação de um serviço de busca de CNPJ do serviço `cnpj-finder`. O projeto fornece uma estrutura robusta para o desenvolvimento de serviços em Go, incluindo configurações pré-definidas para observabilidade, containerização e desenvolvimento local.

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
