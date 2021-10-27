declare namespace Screen {
  interface Aptitude {
    ocr: (
      callback: (error: Error | undefined, value: OCRResult[]) => void,
      ocrCoordinates: OCRCoordinates,
    ) => void;
  }
  interface OCRCoordinates {
    top: number;
    left: number;
    width: number;
    height: number;
  }
  interface OCRResult {
    confidence: number;
    text: string;
  }
}
