import Env from "../../config/env";

export type OtelProtocol = 'grpc' | 'http';

export class OtelConfig {
  readonly enabled: boolean;
  readonly serviceName: string;
  readonly serviceVersion: string;
  readonly environment: string;
  readonly endpoint: string;
  readonly protocol: OtelProtocol;

  constructor(env:  Env) {
    this.enabled = env.otelEnabled;

    this.serviceName = env.otelServiceName 
    this.serviceVersion = env.otelServiceVersion 
    this.environment = env.otelResourceEnvironment

    if (!env.otelEndpoint) {
      throw new Error('[OTEL] OTEL_EXPORTER_OTLP_ENDPOINT is required');
    }

    this.endpoint = env.otelEndpoint;
    this.protocol = env.otelProtocol as OtelProtocol
  }
}
