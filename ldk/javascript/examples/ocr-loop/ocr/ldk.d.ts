/* eslint-disable */

declare module 'fastestsmallesttextencoderdecoder';
declare const oliveHelps: OliveHelps.Aptitudes;

declare namespace OliveHelps {
  interface Aptitudes {
    screen: Screen;
  }

  interface OCRResult {
    confidence: number;
    text: string;
  }

  interface Screen {
    ocr: (
      boundingBoxType: number,
      callback: (error: Error | undefined, value: OCRResult[]) => void,
    ) => void;
  }
}
