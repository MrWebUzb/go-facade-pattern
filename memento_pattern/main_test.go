package main

import (
	"testing"

	"github.com/MrWebUzb/facade/memento_pattern/editor"
	"github.com/stretchr/testify/assert"
)

func TestEditor(t *testing.T) {
	te := editor.New(editor.NewTextWindow())
	te.Write("Salom Gopher!\n")
	te.Write("Ishlar qalay!")
	te.Save()

	te.Write("Bu matn o'chirilib yuboriladi.")
	te.Undo()

	assert.Equal(t, "Salom Gopher!\nIshlar qalay!", te.Print())
}
