package editor

// TextEditor ...
type TextEditor struct {
	textWindow      *TextWindow
	textWindowState *TextWindowState
}

// New ...
func New(tw *TextWindow) *TextEditor {
	return &TextEditor{textWindow: tw}
}

// Write ...
func (te *TextEditor) Write(text string) {
	te.textWindow.Write(text)
}

// Print ...
func (te *TextEditor) Print() string {
	return te.textWindow.GetText()
}

// Save ...
func (te *TextEditor) Save() {
	te.textWindowState = te.textWindow.Save()
}

// Undo ...
func (te *TextEditor) Undo() {
	te.textWindow.Restore(te.textWindowState)
}
