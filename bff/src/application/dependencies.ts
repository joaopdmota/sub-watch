import Env from './config/env'; 
import LoggerAdapter from '../infra/logger/loggerAdapter';

export function initializeDependencies(envs: Env, logger: LoggerAdapter) {
    logger.info("Dependencies initialized.");
}