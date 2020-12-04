package editor

import "strings"

// TextWindowState ...
type TextWindowState struct {
	text strings.Builder
}

// NewTextWindowState ...
func NewTextWindowState(s string) *TextWindowState {
	tws := &TextWindowState{strings.Builder{}}
	tws.text.WriteString(s)

	return tws
}

// GetText ...
func (tws TextWindowState) GetText() string {
	return tws.text.String()
}
