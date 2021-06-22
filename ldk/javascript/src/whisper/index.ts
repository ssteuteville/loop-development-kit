/* eslint-disable @typescript-eslint/no-use-before-define */
// import { promisifyMappedBothWithParams } from '../promisify';
// import { mapToExternalWhisper, mapToInternalWhisper } from './mapper';

import {
  NewWhisper,
  Whisper,
  Component,
  UpdateWhisper,
  WhisperComponentType,
  BoxChildComponent,
  StateMap,
} from './types';

export * from './types';

export interface WhisperAptitude {
  /**
   * Adds a new whisper to Olive Helps based on the configuration provided.
   * Returns a promise which provides a reference to the newly created whisper
   *
   * @param whisper The configuration for the whisper being created
   */
  create(whisper: NewWhisper): Promise<Whisper>;
}

// export function create(whisper: NewWhisper): Promise<Whisper> {
//   return promisifyMappedBothWithParams(
//     whisper,
//     mapToInternalWhisper,
//     mapToExternalWhisper,
//     oliveHelps.whisper.create,
//   );
// }

export function create(whisperRequest: NewWhisper): Promise<Whisper> {
  const stateMap = getEditableComponentsStateMap(whisperRequest.components);

  return new Promise((resolve, reject) => {
    try {
      oliveHelps.whisper.create(
        mapToInternalWhisperRequest(whisperRequest, stateMap),
        (error: Error | undefined, internalWhisper: OliveHelps.Whisper) => {
          if (error) {
            reject(error);
          }
          const externalWhisper = mapToExternalWhisper(internalWhisper, stateMap);
          resolve(externalWhisper);
        },
      );
    } catch (e) {
      reject(e);
    }
  });
}

const mapToInternalWhisperRequest = (
  whisperRequest: NewWhisper,
  stateMap: StateMap,
): OliveHelps.NewWhisper => ({
  label: whisperRequest.label,
  components: mapToInternalComponents(whisperRequest.components, stateMap),
  onClose: whisperRequest.onClose,
});

const mapToExternalWhisper = (
  whisper: OliveHelps.Whisper,
  stateMap: StateMap,
): Whisper => ({
  id: whisper.id,
  close: whisper.close,
  componentState: stateMap,
  update: (updateWhisper: UpdateWhisper) => {
    whisper.update(mapToInternalUpdateWhisper(updateWhisper, stateMap));
  },
});

const mapToInternalUpdateWhisper = (updateWhisper: UpdateWhisper,   stateMap: StateMap): OliveHelps.UpdateWhisper => ({
  label: updateWhisper.label,
  components: mapToInternalComponents(updateWhisper.components, stateMap),
});

const mapToInternalComponents = (
  components: Array<Component>,
  stateMap: StateMap,
): Array<OliveHelps.Components> => {
  const internalComponents: Array<OliveHelps.Components> = new Array<OliveHelps.Components>();
  components.forEach((component) => {
    internalComponents.push(mapToInternalComponent(component, stateMap));
  });

  return internalComponents;
};

const mapToInternalComponent = (
  component: Component,
  stateMap: StateMap,
): OliveHelps.Components => {
  switch (component.type) {
    case WhisperComponentType.CollapseBox:
      return {
        id: component.id,
        label: component.label,
        open: component.open,
        children: component.children.map((children) => mapToInternalChildComponent(children, stateMap)),
        type: WhisperComponentType.CollapseBox,
      };
    default:
      return mapToInternalChildComponent(component, stateMap);
  }
};

// TODO: add more editable components
const getEditableComponentsStateMap = (components: Array<Component>): StateMap => {
  const componentStateMap = new Map<string, string|boolean>();
  components.forEach((component) => {
    switch (component.type) {
      // ...
      case 'email':
      case 'textInput':
        if (component.id) {
          componentStateMap.set(component.id, '');
        }
        break;
      default:
        break;
    }
  });

  return componentStateMap;
};

const mapToInternalChildComponent = (
  component: BoxChildComponent,
  stateMap: StateMap,
): OliveHelps.ChildComponents => {
  switch (component.type) {
    case WhisperComponentType.Box:
      // eslint-disable-next-line
      const { onClick } = component;
      if (onClick) {
        return {
          id: component.id,
          alignment: 'justifyContent' in component ? component.justifyContent : component.alignment,
          direction: component.direction,
          children: component.children.map((children) => mapToInternalChildComponent(children, stateMap)),
          type: WhisperComponentType.Box,
          onClick: (error, whisper) => {
            onClick(error, mapToExternalWhisper(whisper, stateMap));
          },
        } as OliveHelps.Box;
      }
      return {
        id: component.id,
        alignment: 'justifyContent' in component ? component.justifyContent : component.alignment,
        direction: component.direction,
        children: component.children.map((children) => mapToInternalChildComponent(children, stateMap)),
        type: WhisperComponentType.Box,
      };
    case WhisperComponentType.Button:
      return {
        ...component,
        onClick: (error, whisper) => {
          component.onClick(error, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.Button;
    case WhisperComponentType.Checkbox:
      return {
        ...component,
        onChange: (error, param, whisper) => {
          if (component.id) {
            stateMap.set(component.id, param);
          }
          component.onChange(error, param, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.Checkbox;
    case WhisperComponentType.Email:
      return {
        ...component,
        onChange: (error, param, whisper) => {
          if (component.id) {
            stateMap.set(component.id, param);
          }
          component.onChange(error, param, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.Email;
    case WhisperComponentType.Link: {
      // eslint-disable-next-line
      const { onClick } = component;
      if (onClick) {
        return {
          ...component,
          onClick: (error, whisper) => {
            onClick(error, mapToExternalWhisper(whisper, stateMap));
          },
        } as OliveHelps.Link;
      }
      return component as OliveHelps.Link;
    }
    case WhisperComponentType.Divider:
    case WhisperComponentType.ListPair:
    case WhisperComponentType.Markdown:
    case WhisperComponentType.Message:
      return component;
    case WhisperComponentType.Number:
      return {
        ...component,
        onChange: (error, param, whisper) => {
          if (component.id) {
            stateMap.set(component.id, param);
          }
          component.onChange(error, param, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.NumberInput;
    case WhisperComponentType.Password:
      return {
        ...component,
        onChange: (error, param, whisper) => {
          if (component.id) {
            stateMap.set(component.id, param);
          }
          component.onChange(error, param, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.Password;
    case WhisperComponentType.RadioGroup:
      return {
        ...component,
        onSelect: (error, param, whisper) => {
          if (component.id) {
            stateMap.set(component.id, param);
          }
          component.onSelect(error, param, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.RadioGroup;
    case WhisperComponentType.Select:
      return {
        ...component,
        onSelect: (error, param, whisper) => {
          if (component.id) {
            stateMap.set(component.id, param);
          }
          component.onSelect(error, param, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.Select;
    case WhisperComponentType.Telephone:
      return {
        ...component,
        onChange: (error, param, whisper) => {
          if (component.id) {
            stateMap.set(component.id, param);
          }
          component.onChange(error, param, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.Telephone;
    case WhisperComponentType.TextInput:
      return {
        ...component,
        onChange: (error, param, whisper) => {
          if (component.id) {
            stateMap.set(component.id, param);
          }
          component.onChange(error, param, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.TextInput;
    default:
      throw new Error('Unexpected component type');
  }
};
