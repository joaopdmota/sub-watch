import './infra/otel/otel'
import { initializeDependencies } from "./application/dependencies";
import WebServer from "./infra/http/webServer"; 
import LoggerAdapter from "./infra/logger/loggerAdapter";
import Env from "./application/config/env";

export default class App {
    private webServer: WebServer;
    private envs: Env;
    private logger: LoggerAdapter;

    constructor() {
        this.envs = new Env();
        this.logger = new LoggerAdapter();

        initializeDependencies(this.envs, this.logger); 
        this.webServer = new WebServer(this.envs, this.logger); 
    }

    public async start() {
        this.webServer.setupRoutes(); 
        this.webServer.start();
    }
}