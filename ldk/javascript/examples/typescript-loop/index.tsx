import * as React from 'react';
import '@oliveai/ldk';
import { StateMap } from '@oliveai/ldk/dist/whisper/types';
import * as ReactWhisper from '@oliveai/ldk/dist/whisper/react/renderer';
const ComponentDemo: React.FunctionComponent<Object> = (props) => {
  const textInputId = 'someUniqueId';
  const [validateButtonClicked, updateButtonClicked] = React.useState(false);
  const [validationError, updateValidationError] = React.useState('');
  const onChange = () => {
    /* do nothing */
  };
  const onValidate = (state: StateMap) => {
    updateValidationError(
      state.get(textInputId) !== 'valid' ? "Input should be 'valid'" : undefined,
    );
  };
  return (
    <oh-whisper label="Simple Form" onClose={() => {}}>
      <>
        <oh-text-input
          label="Input 'valid'"
          id={textInputId}
          key={textInputId}
          onChange={onChange}
          validationError={validationError}
        />
        <oh-button
          label={validateButtonClicked ? 'Validate Again' : 'Validate'}
          onClick={(error, whisperProps) => {
            updateButtonClicked(true);
            onValidate(whisperProps.componentState);
          }}
        />
      </>
    </oh-whisper>
  );
};
ReactWhisper.renderNewWhisper(<ComponentDemo />);
