export class Note {
  constructor({
    url,
    text,
  }) {
    this.url = url;
    this.text = text;
  }

  setUrl(val) {
    this.url = val;
  }

  setText(val) {
    this.text = val;
  }
}

export class Notes {
  constructor() {
    this.notes = [];
  }

  save(note) {
    this.notes.push(note);
  }

  get(url) {
    return this.notes.filter((note) => note.url === url)[0];
  }
}

export default Note;
