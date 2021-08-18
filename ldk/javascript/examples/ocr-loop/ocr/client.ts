export const processOcr = (boundingBoxType: number): Promise<OliveHelps.OCRResult[]> =>
  new Promise((resolve, reject) => {
    oliveHelps.screen.ocr(
      boundingBoxType,
      (error: Error | undefined, ocrResults: OliveHelps.OCRResult[]) => {
        if (error) {
          reject(error);
        }
        resolve(ocrResults);
      },
    );
  });
