import { screen } from '@oliveai/ldk';
import { BoundingBoxType } from '@oliveai/ldk/dist/screen';

function clipboardListenAndWhisper() {
  screen.ocr(BoundingBoxType.Block).then((result) => {
    console.log(`result: ${result}`);
  });
}

clipboardListenAndWhisper();
