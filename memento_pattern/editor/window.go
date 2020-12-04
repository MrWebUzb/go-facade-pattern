package editor

import "strings"

// TextWindow ...
type TextWindow struct {
	currentText strings.Builder
}

// NewTextWindow ...
func NewTextWindow() *TextWindow {
	return &TextWindow{
		currentText: strings.Builder{},
	}
}

// Write ...
func (tw *TextWindow) Write(text string) {
	tw.currentText.WriteString(text)
}

// GetText ...
func (tw *TextWindow) GetText() string {
	return tw.currentText.String()
}

// Save ...
func (tw *TextWindow) Save() *TextWindowState {
	return NewTextWindowState(tw.currentText.String())
}

// Restore ...
func (tw *TextWindow) Restore(saved *TextWindowState) {
	var s strings.Builder
	s.WriteString(saved.GetText())
	tw.currentText = s
}
