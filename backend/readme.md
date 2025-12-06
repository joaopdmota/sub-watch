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

## Estrutura do Projeto

O projeto segue uma organização modular (Clean Architecture / Hexagonal):

- `cmd/server`: Ponto de entrada da aplicação (Main).
- `application`: Lógica de negócios (Use Cases) e regras da aplicação.
- `infra`: Implementações técnicas (HTTP Server, Logs, OTel, Banco de Dados).
- `otel-collector-config.yaml`: Configuração do pipeline de observabilidade.
