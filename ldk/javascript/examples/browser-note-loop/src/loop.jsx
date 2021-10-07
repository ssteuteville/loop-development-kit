import * as React from 'react';
import * as Renderer from '@oliveai/ldk/dist/whisper/react/renderer';
import { browser, ui } from '@oliveai/ldk';
import { Note, Notes } from './Note';

// todo: Global in-memory state for the sake of the demo.
const notes = new Notes();

const EditNoteWhisper = (props) => {
  const onChangeURL = (error, value) => props.note.setUrl(value);
  const onChangeText = (error, value) => props.note.setText(value);
  const onClickClose = (error, whisper) => whisper.close();
  const onClickSave = (error, whisper) => {
    notes.save(props.note);
    whisper.close();
  };
  return (
    <oh-whisper label={props.label} onClose={() => {}}>
      <oh-markdown body="Add a note for your webpage here." />
      <oh-text-input label="Webpage URL" onChange={onChangeURL} value={props.note.url? props.note.url : ''} />
      <oh-rich-text-editor onChange={onChangeText} />
      <oh-box direction="horizontal" justifyContent="flex-end">
        <oh-button label="Close" onClick={onClickClose} />
        <oh-button label="Save" onClick={onClickSave} />
      </oh-box>
    </oh-whisper>
  );
};

const WebpageNotesWhisper = (props) => {
  return (
    <oh-whisper label="Webpage Notes" onClose={() => {}}>
      <oh-markdown body={props.note.text? props.note.text : ''} />
      <oh-link text={props.note.url} href={props.note.url} />
      <oh-box direction="horizontal" justifyContent="flex-end">
        <oh-button label="Edit" onClick={() => {}} />
      </oh-box>
    </oh-whisper>
  );
};

export const start = async () => {
  let lastURL = "";

  await browser.listenNavigation(async (details) => {
    if (details.parentFrameId < 0) {
      lastURL = details.url;
      const note = notes.get(lastURL);
      if (note) {
        await Renderer.renderNewWhisper(
          <WebpageNotesWhisper note={note} />
        );
      }
    }
  });

  await ui.listenSearchbar(async (value) => {
    if (value === 'new note') {
      const note = new Note({
        url: lastURL
      });
      await Renderer.renderNewWhisper(
        <EditNoteWhisper note={note} label="New Note" />
      );
    }
  });
};
