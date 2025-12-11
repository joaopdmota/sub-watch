import { Request, Response, NextFunction } from "express";
import {Logger} from "../../../application/config/logger";

export const requestLoggerMiddleware = (logger: Logger) => {
    return (req: Request, res: Response, next: NextFunction) => {
        const start = Date.now();

        res.on("finish", () => {
            const duration = Date.now() - start;
            
            logger.info(
                "incoming request",
                {
                    method: req.method,
                    url: req.url,
                    statusCode: res.statusCode,
                    durationMs: duration,
                }
            );
        });

        next();
    };
};