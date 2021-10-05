// import { promisifyWithParamAfterCallback } from '../promisify';
import { OCRResult, BoundingBoxType } from './types';

export * from './types';

export interface Screen {
  ocr: (
    boundingBoxType: number,
    // callback: (error: Error | undefined, value: OCRResult[]) => void,
  ) => void;
}

export function ocr(
  boundingBoxType: BoundingBoxType,
  //   callback: (error: Error | undefined, value: OCRResult[]) => void,
): Promise<OCRResult[]> {
  return new Promise((resolve, reject) => {
    oliveHelps.screen.ocr(boundingBoxType, (error: Error | undefined, ocrResults: OCRResult[]) => {
      if (error) {
        reject(error);
      }
      resolve(ocrResults);
    });
  });
}
