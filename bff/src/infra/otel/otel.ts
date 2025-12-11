import { NodeSDK } from "@opentelemetry/sdk-node";
import { Resource } from "@opentelemetry/resources";
import { SemanticResourceAttributes } from "@opentelemetry/semantic-conventions";
import { OTLPTraceExporter } from "@opentelemetry/exporter-trace-otlp-grpc";
import { getNodeAutoInstrumentations } from "@opentelemetry/auto-instrumentations-node";
import * as grpc from '@grpc/grpc-js'; 

const OTEL_ENABLED = process.env.OTEL_ENABLED === 'true';
const TRACES_ENDPOINT = process.env.OTEL_EXPORTER_OTLP_TRACES_ENDPOINT || 'grpc://otel-collector:4317';
const SERVICE_NAME = process.env.OTEL_SERVICE_NAME || process.env.SERVICE_NAME || 'default-node-service';
const SERVICE_VERSION = process.env.OTEL_SERVICE_VERSION || '1.0.0';
const ENVIRONMENT = process.env.OTEL_RESOURCE_ENVIRONMENT || 'development';

(() => {
    if (!OTEL_ENABLED) {
        console.log("[OTEL] OpenTelemetry disabled by configuration.");
        return;
    }

    const otelEndpoint = TRACES_ENDPOINT.startsWith('grpc://')
        ? TRACES_ENDPOINT.replace('grpc://', '')
        : TRACES_ENDPOINT.replace(/^https?:\/\//, '');

    const resource = new Resource({
        [SemanticResourceAttributes.SERVICE_NAME]: SERVICE_NAME,
        [SemanticResourceAttributes.SERVICE_VERSION]: SERVICE_VERSION,
        "deployment.environment": ENVIRONMENT,
    });

    const traceExporter = new OTLPTraceExporter({
        url: otelEndpoint,
        credentials: grpc.credentials.createInsecure(), 
    });

    const sdk = new NodeSDK({
        resource,
        traceExporter,
        instrumentations: [getNodeAutoInstrumentations()],
    });

    try {
        sdk.start();
        console.log(
            `[OTEL] OpenTelemetry SDK activated. Sending INSECURE to: ${otelEndpoint}`,
        );
    } catch (err) {
        console.error("[OTEL] Critical failure to start SDK. Tracing disabled.", err);
        return;
    }
    
    process.on("SIGTERM", () => {
        sdk
          .shutdown()
          .then(() => console.log("[OTEL] Sdk shutdown with success (SIGTERM)"))
          .catch((err) =>
            console.error("[OTEL] Error on shutdown SDK (SIGTERM)", err)
          )
          .finally(() => process.exit(0));
    });

})();