import Env from './config/env'; 
import LoggerAdapter from '../infra/loggerAdapter';

export function initializeDependencies(envs: Env, logger: LoggerAdapter) {
    logger.info("Logger e dependências inicializados.");
}