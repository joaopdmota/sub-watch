import express, { Express } from "express";

export default class WebServer {
    private app: Express;
    
    constructor() {
        this.app = express();
        this.app.use(express.json());
    } 

    public getApp(): Express {
        return this.app;
    }
}