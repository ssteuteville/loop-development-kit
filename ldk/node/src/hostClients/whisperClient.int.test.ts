import FileSystemClient from './whisperClient';
import { Logger } from '../logging';

const logger = new Logger('integration-test-logger');

describe('WhisperHostClient', () => {
  let subject: FileSystemClient;

  beforeAll(async () => {
    subject = new FileSystemClient();
    const connInfo = {
      address: 'localhost:4770',
      serviceId: 1,
      network: 'tcp',
    };
    const session = {
      loopid: 'LOOP_ID',
      token: 'TOKEN',
    };

    await subject.connect(connInfo, session, logger);
  });

  describe('#markdownWhisper', () => {
    test('on overall service failure rejects with an error', async () => {
      const request = { label: 'Hello', markdown: '## World' };

      await expect(
        subject.markdownWhisper(request).promise(),
      ).rejects.toBeDefined();
    });
  });
});
