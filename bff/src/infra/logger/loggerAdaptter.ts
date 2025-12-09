import pino from "pino";

import { Logger } from '../../application/logger';

export class PinoLogger implements Logger {
  private readonly logger: pino.Logger;

  constructor(defaultMeta: Record<string, unknown> = {}) {
    this.logger = pino({
      level: 'info',
      base: defaultMeta,
    });
  }

  child(meta: Record<string, unknown>): PinoLogger {
    return new PinoLogger({
      ...meta,
    });
  }

  info(message: string, meta?: Record<string, unknown>): void {
    this.logger.info(meta ?? {}, message);
  }

  error(message: string, meta?: Record<string, unknown>): void {
    this.logger.error(meta ?? {}, message);
  }

  warn(message: string, meta?: Record<string, unknown>): void {
    this.logger.warn(meta ?? {}, message);
  }

  debug(message: string, meta?: Record<string, unknown>): void {
    this.logger.debug(meta ?? {}, message);
  }
}
