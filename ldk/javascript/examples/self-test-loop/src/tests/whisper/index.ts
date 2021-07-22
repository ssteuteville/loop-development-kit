/* eslint-disable no-async-promise-executor */
import { clipboard, whisper, network, keyboard } from '@oliveai/ldk';
import {
  JustifyContent,
  Direction,
  ButtonStyle,
  Urgency,
  WhisperComponentType,
  TextAlign,
  ButtonSize,
  Whisper,
  NewWhisper,
  Component,
  DateTimeType,
  AudioVolume,
} from '@oliveai/ldk/dist/whisper/types';
import { stripIndent } from 'common-tags';
import { resolveRejectButtons } from './utils';

export const testMarkdownWhisper = (): Promise<boolean> =>
  new Promise(async (resolve, reject) => {
    // const options = ['M12.01', 'M00.123'];
    //  const imgUrl = 'data:image/gif;base64,R0lGODlhNgAKAPcDAPn7+Pj69/f59v3+/f7+/vr7+vT38/r7+fb59fz9+/X49Pv8+/z9/P39/P7//v7+/fb59vn7+f39/fv8+vT38vf59fr8+dLgzvn69+vx6ezy6vP38vH28O/07e3y6vL28f///u3z6/L28O7z7PD17vD17/H18PD07vP28vP28fH17+/z7bzRte7z7ff69+zy6+vx6vT39OXt4ujv5urw59vm1+Dq3d7o2+Lr3+bu5OLr4OTs4d/p26vFot3n2uHq3erx6Nzm2OXt48jZwtHfzMzcx7/TuMXXv9Phztnl1tXi0NDeysnZw9jk1Nbj0s3dyM7dycPWva7HpanEodTi0NDfzMrbxMLWvMvbxsrbxdbj0+jv5dzn2efu5Onw5trl1snaxKC+ltPgz8bYwc7dysXXwM3cyKTBm7PLq97p29rm1+Ls4L3St8DUurvQtMbYwMTXvrfOsLLKqr/TubXMrrfNr6bDnrDJqK/Ip7TLrLXMrYWsebnPsrjPsdrm1r3StrrQsouwf8zbxsbXwMLVu8vcxoaservRtOTt4X6nceHr3urw6ODp3dzo2fL18NHgzMfZwqzGo6K/mbrPssvbxajDn4uvfoCoc6XCnHqkbefu5aG/mJO1iJm5jpu6kIiufJm6j4KqdqXBm6zGpKPAmXOfZajEn4ywgK7HppO2if///568lN3o2bLKq5S2iY2xgpCzhJO1h+Xu44Srd+rx6Z28lNjl1Zi5jpy7kp+9lYmufpCzhZW2ioCpdI2xgXaiaXylb5u7kXGeY5a3i5e4jIeue46yg5W3ipGzhnija5i5jXCdYYyxgKLAmpu6kXKfY3WgZmiXWG+dYGeXV2qZWoyxgXukbX+oc32mb2CSUG2bXl+ST3ymb4WreGGTUX2mcPz8/NXj0ezx6fH276G+l4ywgWqZW5q5j5O1iZa3jPz+/K7IpfX59KTBnOnv5wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACH/C05FVFNDQVBFMi4wAwEEAAAh/wt4bXAgZGF0YXhtcP8/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG10YSB4bWxuczp4PSJhZG9iZTpuczptZXRhLyIgeDp4bXB0az0iQWRvYmUgWE1QIENvcmUgNi4wLWMwMDYgNzkuMTY0NzUzLCAyMDIxLzAyLzE1LTExOjUyOjEzICAgICAgICAiPjxyZGY6UkRGIHhtbG5zOnJkZj0iaHR0cDovL3d3dy53Lm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZjphYm91dD0iIiD/eG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1uczpzdFJlZj0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL3NUeXBlL1Jlc291cmNlUmVmIyIgeG1wOkNyZWF0b3JUb29sPSJBZG9iZSBQaG90b3Nob3AgMjIuMyAoTWFjaW50b3NoKSJ4bXBNTTpJbnN0YW5jZUlEPSJ4bXAuaWlkOkRGODI2NTlFMTIwMTFFQjkyRTRCRjIyNTc2RTI0Q0Qi/yB4bXBNTTpEb2N1bWVudElEPSJ4bXAuZGlkOkRGMjY1N0FFMTIwMTFFQjkyRTRCRjIyNTc2RTI0Q0QiPiA8eG1wTU06RGVyaXZlZEZyb20gc3RSZWZpbnN0YW5jZUlEPSJ4bXAuaWlkOkRGODI2NTc3RTEyMDExRUI5MkU0QkYyMjU3NkUyNENEIiBzdFJlOmRvY3VtZW50SUQ9InhtcC5kaWQ6REY4MjY1NzhFMTIwMTFFQjkyRTRCRjIyNTc2RTI0Q0QiLz4gPC9yZGY6RGVzY3JpcHRpb24+IDwvcmRmOlJERj4gPC94OnhtcG1ldGE+IDw/cGFja2V0IGVuZP89InIiPz4B//79/Pv6+fj39vX08/Lx8O/u7ezr6uno5+bl5OPi4eDf3t3c29rZ2NfW1dTT0tHQz87NzMvKycjHxsXEw8LBwL++vby7urm4t7a1tLOysbCvrq2sq6qpqKempaSjoqGgn56dnJuamZeWlZSTkpGQj46NjIuKiYiHhoWEg4KBgH9+fXx7enl4d3Z1dHNycXBvbm1sa2ppaGdmZWRjYmFgX15dXFtaWVhXVlVUU1JRUE9OTUxLSklIR0ZFRENCQUA/Pj08Ozo5ODc2NTQzMjEwLy4t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';
  
    const markdown1 = stripIndent`![](https://dev.w3.org/SVG/tools/svgweb/samples/svg-files/AJ_Digital_Camera.svg)`;
    const markdown2 = stripIndent`![](https://dev.w3.org/SVG/tools/svgweb/samples/svg-files/Steps.svg)`;
    let firstDisplayed = true;
    // const markdown = stripIndent`![Alt Text](${imgUrl})`;

    const createdWhisper = await whisper.create({
      label: 'Markdown whisper Test',
      onClose: () => {
        console.debug('closed');
      },
      components: [
        {
          body: markdown1,
          type: WhisperComponentType.Markdown,
        },
        {
          type: WhisperComponentType.Audio,
          url: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3',
          volume: AudioVolume.Low,
        }
      ],
    });

    // space bar listener
    const hotkeys: keyboard.Hotkey = {
      key: ' '
    };

    await keyboard.listenHotkey(hotkeys, (pressed) => {
      if (pressed) {
        if (firstDisplayed) {
          createdWhisper.update({
            components: [
              {
                body: markdown2,
                type: WhisperComponentType.Markdown,
              },
              {
                type: WhisperComponentType.Audio,
                url: 'https://www.soundjay.com/mechanical/gun-gunshot-01.mp3',
                volume: AudioVolume.Medium,
              }
            ]
          });
          firstDisplayed = false;
        } else {
          createdWhisper.update({
            components: [
              {
                body: markdown1,
                type: WhisperComponentType.Markdown,
              },
              {
                type: WhisperComponentType.Audio,
                url: 'https://www.soundjay.com/mechanical/gun-trigger-click-01.mp3',
                volume: AudioVolume.High,
              }
            ]
          });
          firstDisplayed = true;
        }
      }
    });
  });

export const testClickableWhisper = (): Promise<boolean> =>
  new Promise(async (resolve) => {
    await whisper.create({
      label: 'Internal Link Test',
      onClose: () => {
        console.debug('closed');
      },
      components: [
        {
          body: 'Select Option 5',
          type: WhisperComponentType.Markdown,
        },
        {
          type: WhisperComponentType.Link,
          textAlign: TextAlign.Left,
          onClick: () => {
            console.debug('wrong');
          },
          text: `Option 1`,
          style: Urgency.None,
        },
        {
          type: WhisperComponentType.Link,
          textAlign: TextAlign.Left,
          onClick: () => {
            console.debug('wrong');
          },
          text: `Option 2`,
          style: Urgency.None,
        },
        {
          type: WhisperComponentType.Link,
          textAlign: TextAlign.Left,
          onClick: () => {
            console.debug('wrong');
          },
          text: `Option 3`,
          style: Urgency.None,
        },
        {
          type: WhisperComponentType.Link,
          textAlign: TextAlign.Left,
          onClick: () => {
            console.debug('wrong');
          },
          text: `Option 4`,
          style: Urgency.None,
        },
        {
          type: WhisperComponentType.Link,
          textAlign: TextAlign.Left,
          onClick: (error: Error, onClickWhisper: Whisper) => {
            onClickWhisper.close((e) => console.log(e));
            resolve(true);
          },
          text: `Option 5`,
          style: Urgency.None,
        },
      ],
    });
  });

export const testBoxInBox = (): Promise<boolean> =>
  new Promise(async (resolve, reject) => {
    try {
      await whisper.create({
        label: 'Box in the box',
        onClose: () => {
          console.debug('closed');
        },
        components: [
          {
            type: WhisperComponentType.Markdown,
            body: stripIndent`
              # Box in the Box Example
              `,
          },
          {
            type: WhisperComponentType.Box,
            alignment: JustifyContent.Center,
            direction: Direction.Horizontal,
            children: [
              {
                type: WhisperComponentType.Box,
                alignment: JustifyContent.Left,
                direction: Direction.Vertical,
                children: [
                  {
                    type: WhisperComponentType.Markdown,
                    body: stripIndent`
                      **Header Left**

                      Some text on the left
                      `,
                  },
                  {
                    type: WhisperComponentType.TextInput,
                    label: 'Left Input',
                    onChange: (value) => {
                      console.debug(`Input value changed: ${value}`);
                    },
                  },
                ],
              },
              {
                type: WhisperComponentType.Box,
                alignment: JustifyContent.Right,
                direction: Direction.Vertical,
                children: [
                  {
                    type: WhisperComponentType.Markdown,
                    body: stripIndent`
                      **Header Right**

                      Some text on the right
                      `,
                  },
                  {
                    type: WhisperComponentType.TextInput,
                    label: 'Right Input',
                    onChange: (value) => {
                      console.debug(`Input value changed: ${value}`);
                    },
                  },
                ],
              },
            ],
          },
          resolveRejectButtons(resolve, reject),
        ],
      });
    } catch (error) {
      console.error(error);

      reject(error);
    }
  });

export const testClickableButton = (): Promise<boolean> =>
  new Promise(async (resolve, reject) => {
    await whisper.create({
      label: 'Button Test',
      onClose: () => {
        console.debug('closed');
      },
      components: [
        {
          body: 'Click the correct button',
          type: WhisperComponentType.Markdown,
        },
        {
          alignment: JustifyContent.SpaceEvenly,
          direction: Direction.Horizontal,
          children: [
            {
              buttonStyle: ButtonStyle.Secondary,
              label: `Don't click me`,
              onClick: () => console.debug(`Why'd you do that?`),
              type: WhisperComponentType.Button,
              size: ButtonSize.Large,
            },
            {
              buttonStyle: ButtonStyle.Text,
              label: `Me neither`,
              onClick: () => console.debug(`Why'd you do that?`),
              type: WhisperComponentType.Button,
              size: ButtonSize.Small,
            },
            {
              label: `Click me`,
              onClick: (error: Error, onClickWhisper: Whisper) => {
                onClickWhisper.close((e) => console.error(e));
                resolve(true);
              },
              type: WhisperComponentType.Button,
            },
          ],
          type: WhisperComponentType.Box,
        },
        {
          alignment: JustifyContent.SpaceEvenly,
          direction: Direction.Horizontal,
          children: [
            {
              label: `Disabled Primary`,
              disabled: true,
              onClick: (error: Error, onClickWhisper: Whisper) => {
                onClickWhisper.close((e) => console.error(e));
                reject(new Error(`Shouldn't be able to click disabled button`));
              },
              type: WhisperComponentType.Button,
              size: ButtonSize.Large,
            },
            {
              label: `Disabled Secondary`,
              buttonStyle: ButtonStyle.Secondary,
              disabled: true,
              onClick: (error: Error, onClickWhisper: Whisper) => {
                onClickWhisper.close((w) => console.error(w));
                reject(new Error(`Shouldn't be able to click disabled button`));
              },
              type: WhisperComponentType.Button,
              size: ButtonSize.Large,
            },
            {
              label: `Disabled Text`,
              buttonStyle: ButtonStyle.Text,
              disabled: true,
              onClick: (error: Error, onClickWhisper: Whisper) => {
                onClickWhisper.close((e) => console.error(e));
                reject(new Error(`Shouldn't be able to click disabled button`));
              },
              type: WhisperComponentType.Button,
              size: ButtonSize.Large,
            },
          ],
          type: WhisperComponentType.Box,
        },
      ],
    });
  });

export const testClickableLink = (): Promise<boolean> =>
  new Promise(async (resolve, reject) => {
    try {
      await whisper.create({
        label: 'External Link Test',
        onClose: () => {
          console.debug('closed');
        },
        components: [
          {
            body: 'Click the link below',
            type: WhisperComponentType.Markdown,
          },
          {
            type: WhisperComponentType.Link,
            textAlign: TextAlign.Left,
            href: 'https://www.google.com',
            text: 'https://www.google.com',
            style: Urgency.None,
          },
          resolveRejectButtons(resolve, reject, 'Url opened in browser', 'Url failed to open'),
        ],
      });
    } catch (error) {
      console.error(error);
      reject(error);
    }
  });

export const testListPairWithCopyableValue = (): Promise<boolean> =>
  new Promise(async (resolve, reject) => {
    const copyableText = 'Click me to copy the value text';

    // reset clipboard value
    await clipboard.write('');

    const createdWhisper = await whisper.create({
      label: 'List Pair Test',
      onClose: () => {
        console.debug('closed');
      },
      components: [
        {
          type: WhisperComponentType.ListPair,
          label: 'I am Mr. Label',
          value: copyableText,
          copyable: true,
          style: Urgency.None,
        },
      ],
    });

    setTimeout(async () => {
      const response = await clipboard.read();
      if (response === copyableText) {
        createdWhisper.close(() => {
          // do nothing.
        });
        resolve(true);
      } else {
        reject(new Error('Incorrect value detected'));
      }
    }, 5000);
  });

export const testListPairWithCopyableLabel = (): Promise<boolean> =>
  new Promise(async (resolve, reject) => {
    const copyableText = 'Click me to copy the label text';

    // reset clipboard value
    await clipboard.write('');

    const createdWhisper = await whisper.create({
      label: 'List Pair Test',
      onClose: () => {
        console.debug('closed');
      },
      components: [
        {
          type: WhisperComponentType.ListPair,
          label: copyableText,
          value: 'I am Mr. Value',
          labelCopyable: true,
          copyable: false,
          style: Urgency.None,
        },
      ],
    });

    setTimeout(async () => {
      const response = await clipboard.read();
      if (response === copyableText) {
        createdWhisper.close(() => {
          // do nothing.
        });
        resolve(true);
      } else {
        reject(new Error('Incorrect value detected'));
      }
    }, 5000);
  });

export const testFormComponents = (): Promise<boolean> =>
  new Promise((resolve, reject) => {
    const textInput = 'myTextInput';
    const emailInput = 'myEmailInput';
    const selectInput = 'mySelectInput';
    let form: Whisper;

    const config: NewWhisper = {
      label: 'Form Whisper',
      onClose: () => {
        console.debug('closed');
      },
      components: [
        {
          id: textInput,
          label: `Enter 'a'`,
          onChange: () => {
            // do nothing.
          },
          tooltip: 'a?',
          type: WhisperComponentType.TextInput,
        },
        {
          id: emailInput,
          label: 'Enter a@b',
          onChange: () => {
            // do nothing.
          },
          type: WhisperComponentType.Email,
        },
        {
          id: selectInput,
          label: `Select 'blue'`,
          onSelect: () => {
            // do nothing.
          },
          options: ['red', 'blue'],
          type: WhisperComponentType.Select,
        },
        {
          id: 'mySubmitButton',
          label: 'Submit',
          onClick: (error: Error, onClickWhisper: Whisper) => {
            onClickWhisper.componentState.forEach((value: string | number | boolean, key: string) =>
              console.info(key, value),
            );
            if (
              onClickWhisper.componentState.get(textInput) === 'a' &&
              onClickWhisper.componentState.get(emailInput) === 'a@b' &&
              onClickWhisper.componentState.get(selectInput) === 1
            ) {
              form.close((e) => console.error(e));
              resolve(true);
            } else {
              form.close((e) => console.error(e));
              reject(new Error('Please enter correct form values.'));
            }
          },
          type: WhisperComponentType.Button,
        },
        {
          type: WhisperComponentType.Divider,
        },
        {
          header: 'Form render test. Take no action.',
          type: WhisperComponentType.Message,
        },
        {
          label: `Second Text Input`,
          onChange: () => {
            // do nothing.
          },
          value: 'My Initial Value',
          id: 'myTextInputTwo',
          type: WhisperComponentType.TextInput,
        },
        {
          label: `Select 'blue'`,
          onSelect: () => {
            // do nothing.
          },
          options: ['red', 'blue'],
          id: 'mySelectInputTwo',
          type: WhisperComponentType.Select,
        },
        {
          onSelect: () => {
            // do nothing.
          },
          options: ['option1', 'option2'],
          id: 'myRadioGroup',
          type: WhisperComponentType.RadioGroup,
        },
        {
          label: 'Check a box',
          value: true,
          onChange: () => {
            // do nothing.
          },
          id: 'myCheckbox',
          type: WhisperComponentType.Checkbox,
        },
        {
          label: 'Enter an email',
          onChange: () => {
            // do nothing.
          },
          id: 'myEmail',
          type: WhisperComponentType.Email,
        },
        {
          label: 'Enter a number',
          onChange: () => {
            // do nothing.
          },
          id: 'myNumber',
          type: WhisperComponentType.Number,
        },
        {
          label: 'Enter a password',
          onChange: () => {
            // do nothing.
          },
          id: 'myPassword',
          type: WhisperComponentType.Password,
        },
        {
          label: 'Enter a telephone',
          onChange: () => {
            // do nothing.
          },
          id: 'myTelephone',
          type: WhisperComponentType.Telephone,
        },
        {
          id: 'dummySubmitButton',
          label: 'Dummy Submit',
          onClick: (error: Error, onClickWhisper: Whisper) => {
            onClickWhisper.componentState.forEach((value: string | number | boolean, key: string) =>
              console.info(key, value),
            );
          },
          type: WhisperComponentType.Button,
        },
      ],
    };

    whisper.create(config).then((whisperForm: Whisper) => {
      form = whisperForm;
    });
  });

export const testNumberInputs = (): Promise<boolean> =>
  new Promise(async (resolve, reject) => {
    await whisper.create({
      label: 'Number Test',
      components: [
        {
          type: WhisperComponentType.Number,
          label: 'No min, max 10, step 1',
          max: 10,
          step: 1,
          tooltip: 'A tooltip',
          onChange: (error, newValue) => console.log(`New number: ${newValue}`),
        },
        {
          type: WhisperComponentType.Number,
          label: 'No optional fields',
          onChange: (error, newValue) => console.log(`New number: ${newValue}`),
        },
        {
          type: WhisperComponentType.Number,
          label: 'All optional fields',
          value: 0,
          min: 0,
          max: 10,
          step: 0.1,
          tooltip: 'A tooltip',
          onChange: (error, newValue) => console.log(`New number: ${newValue}`),
        },
        {
          type: WhisperComponentType.Telephone,
          label: 'label',
          onChange: (value) => console.log(`Telephone is changed: ${value}`),
          tooltip: 'tooltip',
          value: '09123456789',
        },
        resolveRejectButtons(resolve, reject),
      ],
      onClose: () => {
        console.log('close');
      },
    });
  });

export const testFloatNumberInputs = (): Promise<boolean> =>
  new Promise(async (resolve) => {
    const createdWhisper = await whisper.create({
      label: 'Number Test',
      components: [
        {
          type: WhisperComponentType.Number,
          label: 'Change to 0.6',
          max: 5.5,
          step: 0.1,
          onChange: (error, newValue) => {
            if (newValue === 0.6) {
              createdWhisper.close((e) => console.error(e));
              resolve(true);
            }
          },
        },
      ],
      onClose: () => {
        console.log('close');
      },
    });
  });

export const testNetworkAndListComponents = (): Promise<boolean> =>
  new Promise(async (resolve, reject) => {
    const url = `https://api.fda.gov/food/enforcement.json?search=report_date:[20210101+TO+20210401]&limit=1`;

    const response = await network.httpRequest({
      url,
      method: 'GET',
    });

    console.debug('Network call succeeded, emitting list whisper', url);
    const decodedValue = await network.decode(response.body);
    const { results } = JSON.parse(decodedValue);
    const [recallItem] = results;

    await whisper.create({
      label: 'Latest FDA Food Recall',
      onClose: () => {
        console.debug('closed');
      },
      components: [
        {
          body: recallItem.product_description,
          header: recallItem.recalling_firm,
          style: Urgency.None,
          type: WhisperComponentType.Message,
        },
        {
          type: WhisperComponentType.Divider,
        },
        {
          copyable: true,
          label: 'Reason',
          style: Urgency.None,
          type: WhisperComponentType.ListPair,
          value: recallItem.reason_for_recall,
        },
        {
          copyable: true,
          label: 'Distribution',
          style: Urgency.None,
          type: WhisperComponentType.ListPair,
          value: recallItem.distribution_pattern,
        },
        {
          copyable: true,
          label: 'Quantity',
          style: Urgency.None,
          type: WhisperComponentType.ListPair,
          value: recallItem.product_quantity,
        },
        {
          copyable: true,
          label: 'Codes',
          style: Urgency.None,
          type: WhisperComponentType.ListPair,
          value: recallItem.code_info,
        },
        {
          label: 'Expand',
          open: false,
          children: [
            {
              copyable: true,
              label: 'Recall Type',
              style: Urgency.None,
              type: WhisperComponentType.ListPair,
              value: recallItem.voluntary_mandated,
            },
            {
              copyable: true,
              label: 'Product type',
              style: Urgency.None,
              type: WhisperComponentType.ListPair,
              value: recallItem.product_type,
            },
            {
              copyable: true,
              label: 'Classification',
              style: Urgency.None,
              type: WhisperComponentType.ListPair,
              value: recallItem.classification,
            },
          ],
          type: WhisperComponentType.CollapseBox,
        },
        resolveRejectButtons(resolve, reject),
      ],
    });
  });

export const testDefaultValueForSelectAndRadio = (): Promise<boolean> =>
  new Promise(async (resolve, reject) => {
    try {
      await whisper.create({
        label: 'Are default values displayed correctly?',
        onClose: () => {
          console.debug('closed');
        },
        components: [
          {
            type: WhisperComponentType.Select,
            label: 'Default value: red',
            options: ['green', 'red', 'blue'],
            onSelect: () => {
              // do nothing.
            },
            selected: 1,
          },
          {
            type: WhisperComponentType.RadioGroup,
            onSelect: () => {
              // do nothing.
            },
            options: ['dog', 'cat', 'snake'],
            selected: 1,
          },
          resolveRejectButtons(resolve, reject),
        ],
      });
    } catch (e) {
      console.error(e);
      reject(e);
    }
  });

export const testTooltips = (): Promise<boolean> =>
  new Promise(async (resolve, reject) => {
    try {
      await whisper.create({
        label: 'Tooltip Whisper',
        onClose: () => {
          console.debug('whisper closed');
        },
        components: [
          {
            type: WhisperComponentType.Markdown,
            body: `Hover to see tooltip`,
            tooltip: 'Tooltip for Markdown',
          },
          {
            type: WhisperComponentType.Message,
            header: 'Message Header',
            body: `Hover to see tooltip`,
            style: Urgency.Success,
            tooltip: 'Tooltip for Message',
            textAlign: TextAlign.Left,
          },
          {
            type: WhisperComponentType.Button,
            label: 'Hover to see tooltip',
            onClick: () => {
              // do nothing.
            },
            tooltip: 'Tooltip for Button',
            disabled: true,
          },
          {
            type: WhisperComponentType.Markdown,
            body: stripIndent`
              # Are all tooltips rendered?
              `,
          },
          resolveRejectButtons(resolve, reject),
        ],
      });
    } catch (error) {
      console.error(error);
      reject(error);
    }
  });

export const testClickableBox = (): Promise<boolean> =>
  new Promise(async (resolve) => {
    await whisper.create({
      label: 'Clickable Box Test',
      onClose: () => {
        console.debug('closed');
      },
      components: [
        {
          type: WhisperComponentType.Markdown,
          body: 'Click the correct button',
        },
        {
          type: WhisperComponentType.Box,
          alignment: JustifyContent.SpaceEvenly,
          direction: Direction.Horizontal,
          onClick: () => {
            console.debug('The toggles...they do nothing');
          },
          children: [
            {
              type: WhisperComponentType.Markdown,
              body: `Don't click me`,
            },
          ],
        },
        {
          type: WhisperComponentType.Box,
          alignment: JustifyContent.SpaceEvenly,
          direction: Direction.Horizontal,
          onClick: (error: Error, onClickWhisper: Whisper) => {
            resolve(true);
            onClickWhisper.close((e) => console.error(e));
          },
          children: [
            {
              body: `Click me`,
              type: WhisperComponentType.Markdown,
            },
          ],
        },
      ],
    });
  });

export const testClickableBoxNestingBoxes = (): Promise<boolean> =>
  new Promise(async (resolve, reject) => {
    await whisper.create({
      label: 'Nested Clickable Box Test',
      onClose: () => {
        console.debug('closed');
      },
      components: [
        {
          body: 'Click the text below',
          type: WhisperComponentType.Markdown,
        },
        {
          alignment: JustifyContent.SpaceEvenly,
          direction: Direction.Horizontal,
          onClick: (error: Error, onClickWhisper: Whisper) => {
            onClickWhisper.close((e) => console.error(e));
            reject(new Error('Outer element clicked'));
          },
          children: [
            {
              alignment: JustifyContent.SpaceEvenly,
              direction: Direction.Horizontal,
              onClick: (error: Error, onClickWhisper: Whisper) => {
                onClickWhisper.close((e) => console.error(e));
                resolve(true);
              },
              children: [
                {
                  body: `Click these words`,
                  type: WhisperComponentType.Markdown,
                },
              ],
              type: WhisperComponentType.Box,
            },
            {
              alignment: JustifyContent.SpaceEvenly,
              direction: Direction.Horizontal,
              children: [
                {
                  body: `Don't click here`,
                  type: WhisperComponentType.Markdown,
                },
              ],
              type: WhisperComponentType.Box,
            },
          ],
          type: WhisperComponentType.Box,
        },
      ],
    });
  });

export const testClickableBoxNestingButtons = (): Promise<boolean> =>
  new Promise(async (resolve, reject) => {
    await whisper.create({
      label: 'Nested Button in Clickable Box ',
      onClose: () => {
        console.debug('closed');
      },
      components: [
        {
          body: 'Click the button',
          type: WhisperComponentType.Markdown,
        },
        {
          alignment: JustifyContent.SpaceEvenly,
          direction: Direction.Horizontal,
          onClick: (error: Error, onClickWhisper: Whisper) => {
            onClickWhisper.close((e) => console.error(e));
            reject(new Error('Outer element clicked'));
          },
          children: [
            {
              label: `Click me`,
              onClick: (error: Error, onClickWhisper: Whisper) => {
                onClickWhisper.close((e) => console.error(e));
                resolve(true);
              },
              type: WhisperComponentType.Button,
            },
          ],
          type: WhisperComponentType.Box,
        },
      ],
    });
  });

export const testClickableBoxNestingLinks = (): Promise<boolean> =>
  new Promise(async (resolve, reject) => {
    await whisper.create({
      label: 'Nested Links in Clickable Box ',
      onClose: () => {
        console.debug('closed');
      },
      components: [
        {
          body: 'Click the link',
          type: WhisperComponentType.Markdown,
        },
        {
          alignment: JustifyContent.SpaceEvenly,
          direction: Direction.Horizontal,
          onClick: (error: Error, onClickWhisper: Whisper) => {
            onClickWhisper.close((e) => console.error(e));
            reject(new Error('Outer element clicked'));
          },
          children: [
            {
              type: WhisperComponentType.Link,
              textAlign: TextAlign.Left,
              onClick: (error: Error, onClickWhisper: Whisper) => {
                onClickWhisper.close((e) => console.log(e));
                resolve(true);
              },
              text: `Click this link`,
              style: Urgency.None,
            },
          ],
          type: WhisperComponentType.Box,
        },
      ],
    });
  });

const areAllResolved = (resolverMap: Map<string, boolean>) => {
  let result = true;
  resolverMap.forEach((value) => {
    if (!value) {
      result = false;
    }
  });

  return result;
};

const onActionWrapper = (
  error: Error,
  actionType: string,
  resolverMap: Map<string, boolean>,
  createdWhisper: Whisper,
  resolve: (value: boolean) => void,
  reject: (reason?: Error) => void,
) => {
  if (error) {
    console.error(error);
    reject(error);
  }
  console.debug(`Received ${actionType} event`);
  resolverMap.set(actionType, true);

  if (areAllResolved(resolverMap)) {
    resolve(true);
    createdWhisper.close(() => {
      // do nothing.
    });
  }
};

export const testOnBlurAndOnFocus = (): Promise<boolean> =>
  new Promise(async (resolve, reject) => {
    const resolverMap = new Map([
      ['BlurText', false],
      ['FocusText', false],
      ['BlurNumber', false],
      ['FocusNumber', false],
      ['BlurTelephone', false],
      ['FocusTelephone', false],
      ['BlurPassword', false],
      ['FocusPassword', false],
      ['BlurEmail', false],
      ['FocusEmail', false],
    ]);

    try {
      const createdWhisper = await whisper.create({
        label: 'OnBlur Whisper',
        onClose: () => {
          console.debug('whisper closed');
        },
        components: [
          {
            type: WhisperComponentType.TextInput,
            label: 'Text onBlur',
            onChange: () => {
              // do nothing.
            },
            onBlur: (error: Error) => {
              onActionWrapper(error, 'BlurText', resolverMap, createdWhisper, resolve, reject);
            },
            onFocus: (error: Error) => {
              onActionWrapper(error, 'FocusText', resolverMap, createdWhisper, resolve, reject);
            },
          },
          {
            type: WhisperComponentType.Telephone,
            label: 'Telephone onBlur',
            onChange: () => {
              // do nothing.
            },
            onBlur: (error) => {
              onActionWrapper(error, 'BlurTelephone', resolverMap, createdWhisper, resolve, reject);
            },
            onFocus: (error: Error) => {
              onActionWrapper(
                error,
                'FocusTelephone',
                resolverMap,
                createdWhisper,
                resolve,
                reject,
              );
            },
          },
          {
            type: WhisperComponentType.Email,
            label: 'Email onBlur',
            onChange: () => {
              // do nothing.
            },
            onBlur: (error) => {
              onActionWrapper(error, 'BlurEmail', resolverMap, createdWhisper, resolve, reject);
            },
            onFocus: (error: Error) => {
              onActionWrapper(error, 'FocusEmail', resolverMap, createdWhisper, resolve, reject);
            },
          },
          {
            type: WhisperComponentType.Number,
            label: 'Number onBlur',
            onChange: () => {
              // do nothing.
            },
            onBlur: (error) => {
              onActionWrapper(error, 'BlurNumber', resolverMap, createdWhisper, resolve, reject);
            },
            onFocus: (error: Error) => {
              onActionWrapper(error, 'FocusNumber', resolverMap, createdWhisper, resolve, reject);
            },
          },
          {
            type: WhisperComponentType.Password,
            label: 'Password onBlur',
            onChange: () => {
              // do nothing.
            },
            onBlur: (error) => {
              onActionWrapper(error, 'BlurPassword', resolverMap, createdWhisper, resolve, reject);
            },
            onFocus: (error: Error) => {
              onActionWrapper(error, 'FocusPassword', resolverMap, createdWhisper, resolve, reject);
            },
          },
        ],
      });
    } catch (error) {
      console.error(error);
      reject(error);
    }
  });

export const testCollapseBoxOnClick = (): Promise<boolean> =>
  new Promise(async (resolve, reject) => {
    const resolutions = {
      expand: false,
      collapse: false,
    };
    const bothResolved = () => resolutions.expand && resolutions.collapse;
    await whisper.create({
      label: 'Expand / Collapse OnClick Callback Test',
      components: [
        {
          type: WhisperComponentType.Markdown,
          body: 'Markdown outside CollapseBox',
        },
        {
          type: WhisperComponentType.Divider,
        },
        {
          type: WhisperComponentType.Markdown,
          body: 'Markdown outside CollapseBox',
        },
        {
          type: WhisperComponentType.CollapseBox,
          open: false,
          label: 'Expand me!',
          onClick: (error: Error, open: boolean, onClickWhisper: Whisper) => {
            if (open) {
              resolutions.expand = true;
              if (bothResolved()) {
                onClickWhisper.close((e) => console.error(e));
                resolve(true);
              }
            } else {
              if (resolutions.expand) {
                return;
              }
              reject(new Error('CollapseBox should have reported open as true'));
            }
          },
          children: [
            {
              type: WhisperComponentType.Markdown,
              body: 'Good job!',
            },
            {
              type: WhisperComponentType.Divider,
            },
            {
              type: WhisperComponentType.Markdown,
              body: 'Good job!',
            },
            {
              type: WhisperComponentType.Divider,
            },
          ],
        },
        {
          type: WhisperComponentType.CollapseBox,
          open: true,
          label: 'Collapse me!',
          onClick: (error: Error, open: boolean, onClickWhisper: Whisper) => {
            if (!open) {
              resolutions.collapse = true;
              if (bothResolved()) {
                onClickWhisper.close((e) => console.error(e));
                resolve(true);
              }
            } else {
              if (resolutions.collapse) {
                return;
              }
              reject(new Error('CollapseBox should have reported open as false'));
            }
          },
          children: [
            {
              type: WhisperComponentType.Markdown,
              body: 'Being hidden is my destiny!',
            },
          ],
        },
      ],
    });
  });

export const testDateTime = (): Promise<boolean> =>
  new Promise(async (resolve, reject) => {
    const resolverMap = new Map([
      ['Date', false],
      ['Time', false],
      ['DateTime', false],
    ]);

    try {
      const components: Component[] = [
        {
          type: WhisperComponentType.DateTimeInput,
          key: 'dateId',
          id: 'dateId',
          label: 'Date',
          dateTimeType: DateTimeType.Date,
          onChange: (error: Error, param: string, onChangeWhisper: Whisper) => {
            if (param) {
              console.debug(`Date picker value received: ${param}`);
              onActionWrapper(error, 'Date', resolverMap, onChangeWhisper, resolve, reject);
            }
          },
          tooltip: 'Date picker',
          min: new Date(2020, 0, 1),
          value: new Date(2021, 0, 1),
          max: new Date(2022, 11, 31),
        },
        {
          type: WhisperComponentType.DateTimeInput,
          key: 'timeId',
          id: 'timeId',
          label: 'Time',
          dateTimeType: DateTimeType.Time,
          onChange: (error: Error, param: string, onChangeWhisper: Whisper) => {
            if (param) {
              console.debug(`Time picker value received: ${param}`);
              onActionWrapper(error, 'Time', resolverMap, onChangeWhisper, resolve, reject);
            }
          },
          tooltip: 'Time picker',
          value: new Date(0, 0, 0, 14, 30),
        },
        {
          type: WhisperComponentType.DateTimeInput,
          key: 'dateTimeId',
          id: 'dateTimeId',
          label: 'Date and Time',
          dateTimeType: DateTimeType.DateTime,
          onChange: (error: Error, param: string, onChangeWhisper: Whisper) => {
            if (param) {
              console.debug(`DateTime picker value received: ${param}`);
              onActionWrapper(error, 'DateTime', resolverMap, onChangeWhisper, resolve, reject);
            }
          },
          tooltip: 'Date/Time picker',
          min: new Date(2020, 0, 1),
          value: new Date(2021, 4, 5),
          max: new Date(2022, 11, 31),
        },
      ];
      await whisper.create({
        label: 'Pick date and time',
        components: [
          ...components,
          {
            type: WhisperComponentType.Button,
            label: 'Update',
            onClick: (error: Error, onClickWhisper: Whisper) => {
              onClickWhisper.update({
                components,
              });
            },
          },
        ],
      });
    } catch (error) {
      console.error(error);
      reject(error);
    }
  });
