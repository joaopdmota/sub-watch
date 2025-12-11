// src/config/Config.ts

import * as dotenv from 'dotenv';

// --- Interfaces ---
export interface ServerConfig { PORT: number; NAME: string; }
export interface OtelConfig {
  ENABLED: boolean;
  TRACES_ENDPOINT: string;
  PROTOCOL: string;
  SERVICE_NAME: string;
  SERVICE_VERSION: string;
  ENVIRONMENT: string;
}

export interface AppConfig {
  SERVER: ServerConfig;
  OTEL: OtelConfig;
}

const getNumber = (key: string, defaultValue: number): number => {
  const value = process.env[key];
  return value && !isNaN(Number(value)) ? Number(value) : defaultValue;
};

const getString = (key: string, defaultValue: string): string => {
  return process.env[key] || defaultValue;
};

const getBoolean = (key: string, defaultValue: boolean = false): boolean => {
  const value = process.env[key];
  return typeof value === 'string' ? value.toLowerCase() === 'true' : defaultValue;
};

export default class Config {
    public readonly SERVER: ServerConfig;
    public readonly OTEL: OtelConfig;

    constructor() {
        dotenv.config();

        this.SERVER = {
            PORT: getNumber('SERVICE_PORT', 3001),
            NAME: getString('SERVICE_NAME', 'sub-watch-bff'),
        };

        this.OTEL = {
            ENABLED: getBoolean('OTEL_ENABLED', false),
            TRACES_ENDPOINT: getString(
                'OTEL_EXPORTER_OTLP_TRACES_ENDPOINT',
                getString('OTEL_EXPORTER_OTLP_ENDPOINT', 'http://localhost:4317')
            ),
            PROTOCOL: getString('OTEL_EXPORTER_OTLP_PROTOCOL', 'grpc'),
            SERVICE_NAME: getString('OTEL_SERVICE_NAME', 'sub-watch-bff'),
            SERVICE_VERSION: getString('OTEL_SERVICE_VERSION', '1.0.0'),
            ENVIRONMENT: getString('OTEL_RESOURCE_ENVIRONMENT', 'local'),
        };
    }
}