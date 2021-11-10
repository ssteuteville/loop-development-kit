/* eslint-disable prefer-const */
/* eslint-disable camelcase */
/* eslint-disable prefer-destructuring */
/* eslint-disable no-continue */
/* eslint-disable no-restricted-syntax */
import { screen, whisper, window } from '@oliveai/ldk';
import { write } from '../../dist/clipboard';


// export const writeWhisper = (): Promise<boolean> =>
//   new Promise(async (resolve, reject) => {
//     try {

//       whisper.create({
//         label: formLabel,
//         onClose: () => {
//           // do nothing.
//         },
//         components: [...components,
//             resolveRejectButtons(resolve, reject),],
//       });
//     } catch (error) {
//       console.error(error);
//     }
//   });

const writeWhisper = (labelV: string, body: string) =>
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
      {
        type: whisper.WhisperComponentType.Button,
        label: 'Perform OCR',
        onClick: (error, incomingWhisper) => {
          incomingWhisper.close((e) => console.error(e));
          performOcr();
        },
      },
    ],
  });

// eslint-disable-next-line camelcase
const rebuild_image = (results: screen.OCRResult[]) => {
  const lines = [];
  for (const box of results) {
    if (box.level === undefined) {
      continue;
    }
    let text = box.text;
    let cur_line = box.line_num;
    let cur_word = box.word_num;
    let par_num = box.par_num;

    while (lines.length <= par_num) {
      lines.push([]);
    }
    while (lines[par_num].length <= cur_line) {
      lines[par_num].push([]);
    }
    while (lines[par_num][cur_line].length <= cur_word) {
      lines[par_num][cur_line].push('');
    }
    lines[par_num][cur_line][cur_word] = text;
  }

  let full_text = [];
  for (const para: any of lines) {
    let para_temp = [];
    for (const list_of_words of para) {
      para_temp.push(list_of_words.join(' '));
    }
    full_text.push(para_temp.join('\n'));
  }

  return full_text.join('\n\n');
};

function sleep(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

async function performOcr() {
  await sleep(3000); // sleeping for 2s to switch tabs
  window.activeWindow().then((windowInfo) => {
    const ocrCoordinates = {
      top: windowInfo.y,
      left: windowInfo.x,
      width: windowInfo.width,
      height: windowInfo.height,
    };
    console.log('got active window coordinates:');
    console.log(ocrCoordinates);
    console.log('performing ocr with coordinates...');

    screen
      .ocr(ocrCoordinates)
      .then((result) => {
        // console.log("OCR Results: ");
        // console.log(JSON.stringify(result));

        console.log(rebuild_image(result));
        let resFilter = result.filter((res) => res.confidence > 75);
        resFilter = resFilter.map((res) => `${res.text}`);
        writeWhisper(`result`, `${resFilter.join(' ')}`);
      })
      .catch((error) => {
        console.log('error: ');
        console.log(error);
      });
  });
}

async function OcrLoop() {
  writeWhisper(`ocr`, `starting ocr app`);
  // perform_ocr();
}

OcrLoop();
console.log(`starting app`);