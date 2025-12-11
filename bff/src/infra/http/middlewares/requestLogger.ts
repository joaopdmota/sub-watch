
import LoggerAdapter from "../../logger/loggerAdapter";
import { HttpNextFunction, HttpRequest, HttpResponse, AbstractHttpMiddleware} from "./requet";

export function createLoggingMiddleware(logger: LoggerAdapter): AbstractHttpMiddleware {
    return (req: HttpRequest, res: HttpResponse, next: HttpNextFunction) => {
        const start = Date.now();

        res.on("finish", () => {
            const duration = Date.now() - start;
            
            logger.info(
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
    };
}