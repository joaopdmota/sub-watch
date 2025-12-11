
import { OTLPTraceExporter } from '@opentelemetry/exporter-trace-otlp-grpc';
import { getNodeAutoInstrumentations } from '@opentelemetry/auto-instrumentations-node';
import { Resource } from '@opentelemetry/resources';
import { SemanticResourceAttributes } from '@opentelemetry/semantic-conventions';
import { NodeSDK } from '@opentelemetry/sdk-node';

const otelExporter = new OTLPTraceExporter({
  // A URL do seu Otel Collector.
  // Se sua aplicação BFF rodar fora do Docker, use localhost.
  // Se rodar dentro da mesma rede Docker, use o nome do serviço: 'otel-collector:4317'
  url: 'http://localhost:4317',
});

const sdk = new NodeSDK({
  // Defina o nome do serviço que aparecerá no Jaeger
  resource: new Resource({
    [SemanticResourceAttributes.SERVICE_NAME]: 'bff-sub-watch',
  }),
  traceExporter: otelExporter,
  // Habilita a instrumentação automática para módulos populares do Node.js
  instrumentations: [getNodeAutoInstrumentations()],
});

// Inicia o SDK e a instrumentação
try {
  sdk.start();
  console.log('OpenTelemetry SDK started successfully.');
} catch (error) {
  console.error('Error starting OpenTelemetry SDK:', error);
}

// Garante que o SDK seja desligado corretamente ao finalizar a aplicação
process.on('SIGTERM', () => {
  sdk.shutdown()
    .then(() => console.log('Tracing terminated.'))
    .catch((error) => console.log('Error terminating tracing', error))
    .finally(() => process.exit(0));
});
