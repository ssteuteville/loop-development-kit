// import { promisifyWithParamAfterCallback } from '../promisify';
import { OCRResult } from './types';

export * from './types';

export interface Screen {
  ocr: (
    // callback: (error: Error | undefined, value: OCRResult[]) => void,
  ) => void;
}

export function ocr(
  //   callback: (error: Error | undefined, value: OCRResult[]) => void,
): Promise<OCRResult[]> {
  return new Promise((resolve, reject) => {
    oliveHelps.screen.ocr((error: Error | undefined, ocrResults: OCRResult[]) => {
      if (error) {
        reject(error);
      }
      resolve(ocrResults);
    });
  });
}
