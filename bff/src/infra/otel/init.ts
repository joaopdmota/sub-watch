import { OtelConfig } from './otelConfig';
import { OtelSDK } from './otelSdk';
import Env from '../../config/env';

export async function startTelemetry(envs: Env): Promise<void> {
  const config = new OtelConfig(envs);
  const sdk = new OtelSDK(config);

  await sdk.start();

  process.on('SIGTERM', async () => {
    await sdk.shutdown();
    process.exit(0);
  });
}
