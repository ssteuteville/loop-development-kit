import { clipboard, whisper } from '@oliveai/ldk';

const whisperConfig = {
  label: 'External Link Test',
  onClose: () => {
    // Does nothing
  },
  components: [
    {
      type: 'box',
      justifyContent: 'space-between',
      direction: 'horizontal',
      children: [
        {
          type: 'box',
          justifyContent: 'left',
          direction: 'horizontal',
          children: [
            {
              type: 'icon',
              name: 'navigate_before',
              size: 'medium',
              color: 'black',
            },
          ],
        },
        {
          type: 'box',
          justifyContent: 'center',
          direction: 'horizontal',
          children: [
            {
              type: 'message',
              body: 'Telepsychiatry for ian',
              textAlign: 'center',
            },
          ],
          layout: {
            width: 'full',
            flex: 'auto',
          },
        },
      ],
      alignItems: 'center',
    },
    {
      type: 'divider',
    },
    {
      type: 'box',
      justifyContent: 'center',
      direction: 'horizontal',
      children: [
        {
          buttonStyle: 'text',
          type: 'button',
          label: 'More',
        },
      ],
    },
    {
      type: 'box',
      justifyContent: 'left',
      direction: 'horizontal',
      children: [
        {
          type: 'message',
          body: 'cgjfg',
        },
      ],
    },
    {
      type: 'box',
      justifyContent: 'left',
      direction: 'horizontal',
      children: [
        {
          body: '15:24:15',
          type: 'markdown',
        },
      ],
    },
    {
      type: 'box',
      justifyContent: 'right',
      direction: 'horizontal',
      children: [
        {
          type: 'message',
          body: 'dfdf',
        },
      ],
    },
    {
      type: 'box',
      justifyContent: 'right',
      direction: 'horizontal',
      children: [
        {
          body: '15:24:59',
          type: 'markdown',
        },
      ],
    },
    {
      type: 'box',
      justifyContent: 'left',
      direction: 'horizontal',
      children: [
        {
          type: 'message',
          body: 'ghjfg',
        },
      ],
    },
    {
      type: 'box',
      justifyContent: 'left',
      direction: 'horizontal',
      children: [
        {
          body: '15:25:03',
          type: 'markdown',
        },
      ],
    },
    {
      type: 'message',
      body: '15:30:00: You sent a file to Circulo',
      style: 'success',
      textAlign: 'center',
    },
    {
      key: 'message-text-send-chat',
      label: 'Enter message',
      type: 'textInput',
    },
    {
      type: 'box',
      justifyContent: 'left',
      direction: 'horizontal',
      children: [
        {
          buttonStyle: 'primary',
          disabled: true,
          type: 'button',
          label: 'Send',
        },
        {
          href: '#',
          type: 'link',
          text: 'Upload Documents',
        },
      ],
      alignItems: 'center',
    },
  ],
};
whisper.create(whisperConfig).then((whisperObj) => {
  // ...wait for some trigger to occur
  setTimeout(() => {
    whisperObj.update({
      label: 'Link To Documentation',
      components: [
        {
          type: 'box',
          justifyContent: 'space-between',
          direction: 'horizontal',
          children: [
            {
              type: 'box',
              justifyContent: 'left',
              direction: 'horizontal',
              children: [
                {
                  type: 'icon',
                  name: 'navigate_before',
                  size: 'medium',
                  color: 'black',
                },
              ],
            },
            {
              type: 'box',
              justifyContent: 'center',
              direction: 'horizontal',
              children: [
                {
                  type: 'message',
                  body: 'Telepsychiatry for ian',
                  textAlign: 'center',
                },
              ],
              layout: {
                width: 'full',
                flex: 'auto',
              },
            },
          ],
          alignItems: 'center',
        },
        {
          type: 'divider',
        },
        {
          type: 'box',
          justifyContent: 'center',
          direction: 'horizontal',
          children: [
            {
              buttonStyle: 'text',
              type: 'button',
              label: 'More',
            },
          ],
        },
        {
          type: 'box',
          justifyContent: 'left',
          direction: 'horizontal',
          children: [
            {
              type: 'message',
              body: 'fghfg',
            },
          ],
        },
        {
          type: 'box',
          justifyContent: 'left',
          direction: 'horizontal',
          children: [
            {
              body: '15:23:41',
              type: 'markdown',
            },
          ],
        },
        {
          type: 'box',
          justifyContent: 'right',
          direction: 'horizontal',
          children: [
            {
              type: 'message',
              body: 'sfsd',
            },
          ],
        },
        {
          type: 'box',
          justifyContent: 'right',
          direction: 'horizontal',
          children: [
            {
              body: '15:23:46',
              type: 'markdown',
            },
          ],
        },
        {
          type: 'box',
          justifyContent: 'left',
          direction: 'horizontal',
          children: [
            {
              type: 'message',
              body: 'gfhdsfg',
            },
          ],
        },
        {
          type: 'box',
          justifyContent: 'left',
          direction: 'horizontal',
          children: [
            {
              body: '15:23:50',
              type: 'markdown',
            },
          ],
        },
        {
          type: 'box',
          justifyContent: 'right',
          direction: 'horizontal',
          children: [
            {
              type: 'message',
              body: 'fhf',
            },
          ],
        },
        {
          type: 'box',
          justifyContent: 'right',
          direction: 'horizontal',
          children: [
            {
              body: '15:24:11',
              type: 'markdown',
            },
          ],
        },
        {
          type: 'box',
          justifyContent: 'center',
          direction: 'horizontal',
          children: [
            {
              buttonStyle: 'text',
              type: 'button',
              label: 'More',
            },
          ],
        },
        {
          key: 'message-text-send-chat',
          label: 'Enter message',
          type: 'textInput',
        },
        {
          type: 'box',
          justifyContent: 'left',
          direction: 'horizontal',
          children: [
            {
              buttonStyle: 'primary',
              disabled: true,
              type: 'button',
              label: 'Send',
            },
            {
              href: '#',
              type: 'link',
              text: 'Upload Documents',
            },
          ],
          alignItems: 'center',
        },
      ],
    });
  }, 1500);
});
