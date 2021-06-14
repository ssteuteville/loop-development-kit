import { whisper } from '@oliveai/ldk';

const { Markdown, CollapseBox, Divider, Box } = whisper.WhisperComponentType;

export const showFormatTestingExample = async () => {
  await whisper.create({
    label: 'Format Test',
    components: [
      {
        label: 'this is box',
        alignment: whisper.Alignment.Center,
        direction: whisper.Direction.Vertical,
        children: [
          {
            body: 'box1',
            type: Markdown,
          },
          {
            type: Divider,
          },
          {
            body: 'box2',
            type: Markdown,
          },
          {
            type: Divider,
          },
          {
            body: 'box',
            type: Markdown,
          },
          {
            type: Divider,
          },
          {
            body: 'box4',
            type: Markdown,
          },
          {
            type: Divider,
          },
          {
            body: 'box5',
            type: Markdown,
          },
          {
            type: Divider,
          },
          {
            body: 'box6',
            type: Markdown,
          },
          {
            type: Divider,
          },
          {
            body: 'box7',
            type: Markdown,
          },
        ],
        type: Box,
      },
      {
        body: 'markdown',
        type: Markdown,
      },
      {
        type: Divider,
      },
      {
        body: 'markdown1',
        type: Markdown,
      },
      {
        type: Divider,
      },
      {
        body: 'markdown2',
        type: Markdown,
      },
      {
        type: Divider,
      },
      {
        body: 'markdown3',
        type: Markdown,
      },
      {
        type: Divider,
      },
      {
        body: 'markdown4',
        type: Markdown,
      },
      {
        type: Divider,
      },
      {
        label: 'More',
        open: false,
        children: [
          {
            body: 'markdown inside collapsebox',
            type: Markdown,
          },
          {
            type: Divider,
          },
          {
            body: 'markdown inside collapsebo',
            type: Markdown,
          },
          {
            type: Divider,
          },
          {
            body: 'markdown inside collapsebo',
            type: Markdown,
          },
          {
            type: Divider,
          },
          {
            body: 'markdown inside collapsebo',
            type: Markdown,
          },
          {
            type: Divider,
          },
          {
            body: 'markdown inside collapsebo',
            type: Markdown,
          },
          {
            type: Divider,
          },
          {
            body: 'markdown inside collapsebo',
            type: Markdown,
          },
          {
            type: Divider,
          },
          {
            body: 'markdown inside collapsebo',
            type: Markdown,
          },
        ],
        type: CollapseBox,
      },
      {
        body: `this is the end of collapseBox`,
        type: Markdown,
      },
      {
        type: Divider,
      },
    ],
  });
};
