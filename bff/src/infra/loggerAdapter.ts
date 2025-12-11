import pino from "pino";
import { Logger } from "../application/config/logger";

export default class LoggerAdapter implements Logger {
    private logger: pino.Logger;

    constructor() {
        this.logger = pino();
    }

    info(message: string, data?: any): void {
        this.logger.info(data, message);
    }

    error(message: string, data?: any): void {
        this.logger.error(data, message);
    }

    warn(message: string, data?: any): void {
        this.logger.warn(data, message);
    }

    debug(message: string, data?: any): void {
        this.logger.debug(data, message);
    }
}
