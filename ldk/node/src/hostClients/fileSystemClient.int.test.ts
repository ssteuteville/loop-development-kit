import { FileSystemClient } from './fileSystemClient';
import { Logger } from '../logging';

const logger = new Logger('integration-test-logger');

describe('FileSystemHostClient', () => {
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

  describe('#queryFile', () => {
    test('on overall service failure rejects with an error', async () => {
      const request = { file: '/etc/passwd' };

      await expect(subject.queryFile(request)).rejects.toBeDefined();
    });
  });

  describe('#streamFile', () => {
    test('on overall service failure calls back with an error', (done) => {
      const request = {
        file: '/etc/passwd',
      };

      subject.streamFile(request, (error) => {
        expect(error).toBeDefined();
        done();
      });
    });
  });
});
