export interface OCRResult {
  confidence: number;
  text: string;
}

export enum BoundingBoxType {
  Block,
  Paragraph,
  Line,
  Word,
  Symbol,
}
