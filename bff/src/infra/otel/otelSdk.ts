import { NodeSDK } from '@opentelemetry/sdk-node';
import { Resource } from '@opentelemetry/resources';
import { SemanticResourceAttributes } from '@opentelemetry/semantic-conventions';
import { getNodeAutoInstrumentations } from '@opentelemetry/auto-instrumentations-node';
import { OTLPTraceExporter } from '@opentelemetry/exporter-trace-otlp-grpc';

import { OtelConfig } from './otelConfig';

export class OtelSDK {
  private readonly sdk: NodeSDK;

  constructor(private readonly config: OtelConfig) {
    const resource = new Resource({
      [SemanticResourceAttributes.SERVICE_NAME]: config.serviceName,
      [SemanticResourceAttributes.SERVICE_VERSION]: config.serviceVersion,
      'service.environment': config.environment,
    });

    const traceExporter = new OTLPTraceExporter({
      url: config.endpoint,
    });

    this.sdk = new NodeSDK({
      resource,
      traceExporter,
      instrumentations: [getNodeAutoInstrumentations()],
    });
  }

  async start(): Promise<void> {
    if (!this.config.enabled) {
      console.log('[OTEL] Telemetry disabled');
      return;
    }

    await this.sdk.start();
    console.log('[OTEL] OpenTelemetry started:', this.config.serviceName);
  }

  async shutdown(): Promise<void> {
    if (!this.config.enabled) return;
    await this.sdk.shutdown();
    console.log('[OTEL] OpenTelemetry shut down');
  }
}
