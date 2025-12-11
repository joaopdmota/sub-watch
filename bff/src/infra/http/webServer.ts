import express, { Express, NextFunction, Request, Response } from "express";
import LoggerAdapter from "../logger/loggerAdapter";
import Env from "../../application/config/env";

export default class WebServer {
    private app: Express;
    private envs: Env;
    private logger: LoggerAdapter;
    
    constructor(envs: Env, logger: LoggerAdapter) {
        this.app = express();
        this.envs = envs;
        this.logger = logger;
        this.app.use(express.json());

        this.setupMiddleware();
    } 

    private setupMiddleware() {
        this.app.use((req: Request, res: Response, next: NextFunction) => {
            const start = Date.now();

            res.on("finish", () => {
                const duration = Date.now() - start;
                this.logger.info(
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
    }

    public setupRoutes() {
        this.app.get('/', (req: Request, res: Response) => {
            res.send('Ok');
        });

        this.app.get('/hello-world', (req: Request, res: Response) => {
            res.send('Hello World!');
        });
    }

    public start() {
        this.app.listen(this.envs.SERVER.PORT, () => {
            this.logger.info(
                `Serviço ${this.envs.SERVER.NAME} ouvindo na porta ${this.envs.SERVER.PORT}`,
                { port: this.envs.SERVER.PORT, serviceName: this.envs.SERVER.NAME }
            );
        });
    }
}