import { emitOcrRequestWhisper } from './whisperManager';

const main = async () => {
  await emitOcrRequestWhisper();
};

main().catch((e) => console.error(e));
