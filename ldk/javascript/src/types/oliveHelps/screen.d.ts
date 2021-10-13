declare namespace Screen {
  interface Aptitude {
    ocr: (
      callback: (error: Error | undefined, value: OCRResult[]) => void,
    ) => void;
  }

  interface OCRResult {
    confidence: number;
    text: string;
  }
}
