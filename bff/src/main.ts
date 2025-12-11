import './infra/otel/otel';
import { initializeDependencies } from "./application/dependencies";
import LoggerAdapter from "./infra/loggerAdapter";
import Env from "./application/config/env";
import { NextFunction, Request, Response } from "express"; 

const envs = new Env();
const loggerAdapter = new LoggerAdapter();

initializeDependencies(envs, loggerAdapter); 

import WebServer from "./infra/http/webServer"; 
const webServer = new WebServer();
const expressApp = webServer.getApp();

expressApp.use((req: Request, res: Response, next: NextFunction) => {
    const start = Date.now();

    res.on("finish", () => {
        const duration = Date.now() - start;
        loggerAdapter.info(
            `Request: ${req.method} ${req.url}`,
            {
                statusCode: res.statusCode,
                durationMs: duration,
                method: req.method,
                url: req.url
            }
        );
    });

    next();
});

expressApp.get('/', (req: Request, res: Response) => {
    res.send('Hello World!');
});

expressApp.listen(envs.SERVER.PORT, () => {
    loggerAdapter.info(
        `Serviço ${envs.SERVER.NAME} ouvindo na porta ${envs.SERVER.PORT}`,
        { port: envs.SERVER.PORT, serviceName: envs.SERVER.NAME }
    );
});
