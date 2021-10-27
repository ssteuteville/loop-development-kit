export interface OCRCoordinates {
  top: number;
  left: number;
  width: number;
  height: number;
}
export interface OCRResult {
  confidence: number;
  text: string;
}
