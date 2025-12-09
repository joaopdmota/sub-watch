import express, { NextFunction, Request, Response } from "express";
import HttpServer from "./server";
import { Logger } from "../../application/logger";

function expressLogger(logger: Logger) {
  return function (req: Request, res: Response, next: NextFunction) {
    const start = Date.now();

    res.on("finish", () => {
      const duration = Date.now() - start;
      logger.info(
        `${req.method} ${req.originalUrl} - ${res.statusCode} - ${duration}ms`
      );
    });

    next();
  };
}

export default class ExpressAdapter implements HttpServer {
	app: any;

	constructor (readonly logger: Logger) {
		this.app = express();
		this.app.use(express.json());
		this.app.use(expressLogger(this.logger));
	}

	on(method: string, url: string, callback: Function): void {
		this.app[method](url, async function (req: any, res: any) {
			const output = await callback(req.params, req.body, req.headers);
			res.json(output);
		});
	}

	listen(port: number): void {
		this.logger.info(`Server running on port ${port}`);
		this.app.listen(port);
	}

}