import * as guide from './guide';
import * as hotkeyListener from './hotkeyListener';
import * as searchBarListener from './searchBarListener';
import * as formatTesting from './formatTesting';

formatTesting.showFormatTestingExample();
guide.emitInstructionsWhisper();
hotkeyListener.start();
searchBarListener.start();
