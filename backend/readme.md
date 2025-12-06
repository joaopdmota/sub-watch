Este repositório contém uma configuração do Docker Compose para executar um coletor OpenTelemetry, um servidor Zipkin e um servidor de aplicação localmente. Abaixo você encontra uma descrição de cada serviço, suas funcionalidades e como utilizá-los.

## Serviços

### 1. **otel-collector**
O OpenTelemetry Collector é responsável por coletar, processar e exportar dados de telemetria, como métricas e traces.

- **Imagem utilizada**: `otel/opentelemetry-collector-contrib`
- **Volumes**:
  - `./otel-collector-config.yaml:/etc/otelcol-contrib/config.yaml`: Mapeia o arquivo de configuração do coletor.
- **Portas expostas**:
  - `1888`: Para consultas de dados de métricas.
  - `8888`: Porta para o endpoint de debug.
  - `8889`: Endpoint de verificação de estado.
  - `13133`: Health Check.
  - `4317`: Protocolo OTLP gRPC.
  - `4318`: Protocolo OTLP HTTP.
  - `55679`: Porta legacy para OTLP.
- **Dependências**: O serviço depende do `zipkin` para funcionar corretamente.
- **Rede**: Conectado à rede `app_network`.

---

### 2. **zipkin**
O Zipkin é uma ferramenta de rastreamento distribuído que ajuda a depurar sistemas distribuídos.

- **Imagem utilizada**: `openzipkin/zipkin`
- **Portas expostas**:
  - `9411`: Interface web do Zipkin.
- **Variáveis de ambiente**:
  - `STORAGE_TYPE=mem`: Armazena os dados na memória.
  - `ZIPKIN_HTTP_ENABLED=true`: Habilita o servidor HTTP do Zipkin.
  - `ZIPKIN_HTTP_PORT=9411`: Define a porta HTTP do Zipkin.
- **Rede**: Conectado à rede `app_network`.

---

### 3. **server**
Este é o servidor da aplicação, configurado para rodar localmente com suporte a hot reload via [Air](https://github.com/cosmtrek/air).

- **Build**:
  - **Contexto**: Diretório atual (`.`).
  - **Dockerfile**: `Dockerfile.dev`.
- **Portas expostas**:
  - `8080`: Porta principal da aplicação.
  - `8081`: Porta auxiliar (caso utilizada).
- **Variáveis de ambiente**:
  - Carregadas do arquivo `.env`.
- **Volumes**:
  - `.:/app`: Mapeia o diretório do projeto para o container.
  - `/app/tmp`: Cria um volume temporário no container.
- **Comando**:
  - `air`: Realiza o hot reload durante o desenvolvimento.
- **Rede**: Conectado à rede `app_network`.

---

## Rede

- **app_network**:
  - Tipo: `bridge`
  - Responsável por conectar os serviços no mesmo ambiente de rede.

---

## Como Utilizar

1. Certifique-se de ter o [Docker](https://www.docker.com/) e o [Docker Compose](https://docs.docker.com/compose/) instalados no seu ambiente.
2. Crie o arquivo de configuração do OpenTelemetry Collector (`otel-collector-config.yaml`) na raiz do projeto.
3. Crie um arquivo `.env` basedo no arquivo `.env.example` com as variáveis necessárias para o serviço `server`.
4. Execute o comando abaixo para subir os serviços:
   ```bash
   docker-compose up --build
   ```

# Sistema de Consulta de Clima por CEP

Este projeto tem como objetivo desenvolver um sistema em Go que recebe um CEP, identifica a cidade correspondente e retorna as temperaturas atuais em três unidades de medida: Celsius, Fahrenheit e Kelvin.

## Funcionalidade

O sistema possui as seguintes funcionalidades:

- **Receber um CEP válido de 8 dígitos**
- **Pesquisar a localização correspondente ao CEP utilizando a API ViaCEP**
- **Consultar a temperatura atual utilizando a API WeatherAPI**
- **Retornar a temperatura em Celsius, Fahrenheit e Kelvin**
- **Responder adequadamente nos seguintes cenários:**

### Respostas de Sucesso
- **Código HTTP:** 200
- **Response Body:**
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

### Respostas de Erro

### CEP inválido (formato incorreto):
- **Código HTTP:** 422
- **Mensagem:** `invalid zipcode`

### CEP não encontrado:
- **Código HTTP:** 404
- **Mensagem:** `can not find zipcode`

## Requisitos

- O sistema deve validar o formato do CEP e garantir que seja um número válido de 8 dígitos.
- O sistema deve integrar com as APIs ViaCEP e WeatherAPI para obter as informações necessárias.

## Fórmulas de Conversão

### Celsius para Fahrenheit:
\[
F = C * 1.8 + 32
\]

### Celsius para Kelvin:
\[
K = C + 273
\]

## Tecnologias Utilizadas

- **Linguagem:** Go
- **APIs:**
  - ViaCEP (https://viacep.com.br/)
  - WeatherAPI (https://www.weatherapi.com/)
- **Docker:** Para containerização da aplicação
