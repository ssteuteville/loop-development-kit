import { mocked } from 'ts-jest/utils';
import { ClipboardImpl } from './clipboard';

describe('Clipboard', () => {
  let subject: ClipboardImpl;

  beforeEach(() => {
    oliveHelps.clipboard = {
      read: jest.fn(),
      write: jest.fn(),
      listen: jest.fn(),
    };
    subject = new ClipboardImpl();
  });

  describe('read', () => {
    it('returns a promise result with expected clipboard value', () => {
      const expected = 'expected string';
      mocked(oliveHelps.clipboard.read).mockImplementation((cb) => cb(expected));

      const actual = subject.read();

      return expect(actual).resolves.toBe(expected);
    });

    it('returns a rejected promise', () => {
      const exception = 'Exception';
      mocked(oliveHelps.clipboard.read).mockImplementation(() => {
        throw exception;
      });

      const actual = subject.read();

      return expect(actual).rejects.toBe(exception);
    });
  });

  describe('listen', () => {
    it('passed in listen function to olive helps', () => {
      const callback = jest.fn();
      subject.listen(callback);

      expect(oliveHelps.clipboard.listen).toHaveBeenCalledWith(callback);
    });

    it('throws exception when passing in Listen function', () => {
      const exception = 'Exception';
      mocked(oliveHelps.clipboard.listen).mockImplementation(() => {
        throw exception;
      });

      const callback = jest.fn();
      expect(() => subject.listen(callback)).toThrow(exception);
    });
  });

  describe('write', () => {
    it('writes text to an olive helps clipboard', () => {
      const expectedText = 'text';
      mocked(oliveHelps.clipboard.write).mockImplementation((text, cb) => {
        cb();
      });

      const actual = subject.write(expectedText);

      expect(oliveHelps.clipboard.write).toHaveBeenCalledWith(expectedText, expect.any(Function));
      return expect(actual).resolves.toBeUndefined();
    });

    it('returns a rejected promise', () => {
      const exception = 'Exception';
      mocked(oliveHelps.clipboard.write).mockImplementation(() => {
        throw exception;
      });

      const actual = subject.write('text');

      return expect(actual).rejects.toBe(exception);
    });
  });
});
