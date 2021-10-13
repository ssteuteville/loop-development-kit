import { screen, whisper } from '@oliveai/ldk';
import { BoundingBoxType } from '@oliveai/ldk/dist/screen';
import { write } from '../../dist/clipboard';


const writeWhisper = (labelV, body) => {
  whisper.create({
    label: labelV,
    onClose: () => {
      console.log(`Closed Whisper`);
    },
    components: [
      {
        body,
        type: whisper.WhisperComponentType.Markdown,
      },
    ],
  });
};


function OcrLoop() {
  writeWhisper(`ocr`, `starting ocr app`)
  screen.ocr(0).then((result) => {
    console.log("OCR Results: ");
    console.log(JSON.stringify(result));
    writeWhisper(`result`, `${JSON.stringify(result)}`)
  }).catch((error) => {
    console.log("error: ");
    console.log(error);
  });
}

OcrLoop();
console.log(`starting app`)