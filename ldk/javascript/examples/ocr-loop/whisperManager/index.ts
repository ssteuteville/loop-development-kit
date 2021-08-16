import { whisper } from '@oliveai/ldk';
import {
  Direction,
  JustifyContent,
  Whisper,
  WhisperComponentType,
  Select,
  Component,
} from '@oliveai/ldk/dist/whisper/types';
import { processOcr } from '../ocr/client';

export const emitOcrRequestWhisper = async (): Promise<void> => {
  console.info(`Displaying ocr whisper`);
  const boundingBoxTypeId = 'boundingBoxTypeId1';
  try {
    const components: Component[] = [
      {
        type: WhisperComponentType.Select,
        id: boundingBoxTypeId,
        key: boundingBoxTypeId,
        label: 'Select Bounding Box Type',
        options: ['Block', 'Paragraph', 'Line', 'Word', 'Symbol'],
        onSelect: (error: Error) => {
          if (error) {
            console.error(error);

            return;
          }
        },
      },
      {
        type: WhisperComponentType.Box,
        direction: Direction.Horizontal,
        justifyContent: JustifyContent.Right,
        children: [
          {
            type: WhisperComponentType.Button,
            label: 'Perform OCR',
            onClick: async (error: Error, onClickWhisper: Whisper) => {
              if (error) {
                console.error(error);

                return;
              }
              const boundingBoxType = onClickWhisper.componentState.get(
                boundingBoxTypeId,
              ) as number;
              console.info(`Selected boundingBoxType: ${boundingBoxType}`);
              if (boundingBoxType === undefined || boundingBoxType < 0 || boundingBoxType > 4) {
                (components[0] as Select).validationError = 'Please select Bounding Box Type.';
                onClickWhisper.update({ components });
                return;
              }
              (components[0] as Select).validationError = null;
              onClickWhisper.update({
                components: [
                  {
                    type: WhisperComponentType.Markdown,
                    body: `Please see console for results. **Please be patient!** ;)`,
                  },
                  {
                    type: WhisperComponentType.Box,
                    direction: Direction.Horizontal,
                    justifyContent: JustifyContent.Right,
                    children: [
                      {
                        type: WhisperComponentType.Button,
                        label: 'Restart',
                        onClick: (error: Error, onClickWhisper: Whisper) => {
                          if (error) {
                            console.error(error);

                            return;
                          }

                          onClickWhisper.close((error: Error) => {
                            if (error) {
                              console.error(error);

                              return;
                            }
                          });

                          createNewWhisper(components);
                        },
                      },
                    ],
                  },
                ],
              });
              var start = new Date().getTime();
              console.info(`Starting OCR process`);
              const ocrResults = await processOcr(boundingBoxType);
              var end = new Date().getTime();
              console.info(`Ending OCR process`);
              console.info(`OCR process time: ${(end - start)/1000} sec`);

              ocrResults.forEach((ocrResult) => {
                console.info(
                  `Found text: ${ocrResult.text}, with confidence: ${ocrResult.confidence} per your request.`,
                );
              });
            },
          },
        ],
      },
    ];
    createNewWhisper(components);
  } catch (error) {
    console.error(error);
  }
};

const createNewWhisper = async (components: Component[]) => {
  await whisper.create({
    label: 'Please provide OCR criteria',
    onClose: () => {
      // do nothing.
    },
    components,
  });
};
