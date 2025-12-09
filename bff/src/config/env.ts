import { Logger } from "../application/logger"

export default class Env {
  serviceName: string
  servicePort: number
  otelEnabled: boolean
  otelEndpoint: string
  otelServiceName: string
  otelServiceVersion: string
  otelResourceEnvironment: string
  otelProtocol: string
  backendUrl: string
  backendTimeout: number
  logger: Logger

  constructor(logger: Logger) {
    this.logger = logger
    this.backendTimeout = this.envToNumber(process.env.BACKEND_TIMEOUT)
    this.backendUrl = process.env.BACKEND_URL
    this.servicePort = this.envToNumber(process.env.SERVICE_PORT)
    this.serviceName = process.env.SERVICE_NAME
    this.otelEnabled = process.env.OTEL_ENABLED
    this.otelEndpoint = process.env.OTEL_EXPORTER_OTLP_ENDPOINT
    this.otelServiceName = process.env.OTEL_SERVICE_NAME
    this.otelServiceVersion = process.env.OTEL_SERVICE_VERSION
    this.otelResourceEnvironment = process.env.OTEL_RESOURCE_ENVIRONMENT
    this.otelProtocol = process.env.OTEL_EXPORTER_OTLP_PROTOCOL
  }

  envToNumber(value: string): number {
    const numberValue = Number(value)
    if (isNaN(numberValue)) {
      this.logger.error("Invalid number env to convert: " + value)
    }
    return numberValue
  }
}