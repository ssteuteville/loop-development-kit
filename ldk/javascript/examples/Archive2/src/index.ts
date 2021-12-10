import { whisper } from '@oliveai/ldk';

(async function main(): Promise<void> {
  const box: whisper.Box = {
    type: whisper.WhisperComponentType.Box,
    children: [{ type: whisper.WhisperComponentType.Message, body: 'Hello World' }],
    justifyContent: whisper.JustifyContent.Center,
    direction: whisper.Direction.Vertical,
  };
  const link: whisper.Link = {
    type: whisper.WhisperComponentType.Link,
    text: 'Button',
    onClick: (error, thisWhisper) => {
      if (error) console.error(error);
      thisWhisper.update({ label: 'test', components: [box] });
    },
  };
  const link2: whisper.Link = {
    type: whisper.WhisperComponentType.Link,
    text: 'Button 2',
    onClick: (error, thisWhisper) => {
      if (error) console.error(error);
      thisWhisper.update({
        label: 'test',
        components: [{ type: whisper.WhisperComponentType.Markdown, body: '# Hello' }],
      });
    },
  };
  const bread: whisper.Breadcrumbs = {
    type: whisper.WhisperComponentType.Breadcrumbs,
    links: [link2, link, link],
  };
  const md: whisper.Markdown = {
    type: whisper.WhisperComponentType.Markdown,
    body: '![image_name](https://d33wubrfki0l68.cloudfront.net/e7ed9fe4bafe46e275c807d63591f85f9ab246ba/e2d28/assets/images/tux.png)',
  };
  const div: whisper.Divider = { type: whisper.WhisperComponentType.Divider };
  whisper.create({
    label: 'Test',
    onClose: () => {
      console.log('closed');
    },
    components: [md, bread, div, link, link2],
  });


})();
