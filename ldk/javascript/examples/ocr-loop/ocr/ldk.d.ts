/* eslint-disable */

declare module 'fastestsmallesttextencoderdecoder';
declare const oliveHelps: OliveHelps.Aptitudes;

declare namespace OliveHelps {
  interface Aptitudes {
    filesystem: Filesystem;
  }

  interface OCRResult {
    confidence: number;
    text: string;
  }

  interface Filesystem {
    ocr: (
      boundingBoxType: number,
      callback: (error: Error | undefined, value: OCRResult[]) => void,
    ) => void;
  }
}
