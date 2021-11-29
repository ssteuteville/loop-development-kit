import React, { useState, useEffect } from 'react';
import '@oliveai/ldk';
import { StateMap } from '@oliveai/ldk/dist/whisper/types';
import {
  AlignItems,
  Direction,
  JustifyContent,
  StyleSize,
  IconSize,
  Color,
} from '@oliveai/ldk/dist/whisper';
import * as ReactWhisper from '@oliveai/ldk/dist/whisper/react/renderer';

const ComponentDemo: React.FunctionComponent<Object> = (props) => {
  const textInputId = 'someUniqueId';
  const [value, setValue] = React.useState('');

  let t = setTimeout(() => {
    // do nothing
  }, 0);

  return (
    <oh-whisper label="Simple Form" onClose={() => {}}>
      <>
        <oh-box direction={Direction.Vertical} justifyContent={JustifyContent.Left}>
          <oh-text-input
            label="testing"
            onChange={(_, val) => {
              clearTimeout(t);
              t = setTimeout(() => {
                setValue(val);
              }, 500);
            }}
            value={value}
          />
          <oh-button
            label={'Validate Again'}
            onClick={(error, whisperProps) => {
              console.log(`Hello there ${value}`);
            }}
          />
        </oh-box>
      </>
    </oh-whisper>
  );
};
ReactWhisper.renderNewWhisper(<ComponentDemo />);
