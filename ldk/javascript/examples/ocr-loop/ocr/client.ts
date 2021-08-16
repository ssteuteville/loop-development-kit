export const processOcr = (boundingBoxType: number): Promise<OliveHelps.OCRResult[]> => {
    return new Promise((resolve, reject) => {
        oliveHelps.filesystem.ocr(boundingBoxType, (error: Error | undefined, ocrResults: OliveHelps.OCRResult[]) => {
            if (error) {
                reject(error);
            }
            resolve(ocrResults);
        });
    });
}