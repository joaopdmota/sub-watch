import Env from "./config/env";
import ExpressAdapter from "./infra/http/expressAdapter";
import Router from "./infra/http/router";
import { PinoLogger } from "./infra/logger/loggerAdaptter";
import { startTelemetry } from "./infra/otel/init";

async function bootstrap() {
  const logger = new PinoLogger();
  const env = new Env(logger);

  logger.info("Starting application bootstrap...");

  if (env.otelEnabled) {
    try {
      await startTelemetry(env);
      logger.info("Telemetry started successfully");
    } catch (err) {
      logger.error(`Failed to start telemetry: ${(err as Error).message}`);
    }
  }

  const app = new ExpressAdapter(logger);
  new Router(app);

  app.listen(env.servicePort);
}

bootstrap().catch((err) => {
  console.error("Fatal error during bootstrap:", err);
});
