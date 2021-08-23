import {
    promisifyWithParam
} from '../promisify'

export interface Lifecycle {
    onStart(startFunction: () => void): Promise<void>;
}

export function onStart(startFunction: () => void) : Promise<void> {
    return promisifyWithParam(startFunction, oliveHelps.lifecycle.onStart);
}