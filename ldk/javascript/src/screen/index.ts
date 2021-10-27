// import { promisifyWithParamAfterCallback } from '../promisify';
import { OCRResult, OCRCoordinates } from './types';

export * from './types';

export interface Screen {
  ocr: (ocrCoordinates: OCRCoordinates) => void;
}

export function ocr(ocrCoordinates: OCRCoordinates): Promise<OCRResult[]> {
  return new Promise((resolve, reject) => {
    oliveHelps.screen.ocr((error: Error | undefined, ocrResults: OCRResult[]) => {
      if (error) {
        reject(error);
      }
      resolve(ocrResults);
    }, ocrCoordinates);
  });
}
