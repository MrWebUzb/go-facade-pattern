package main

import (
	"fmt"

	"github.com/MrWebUzb/facade/memento_pattern/editor"
)

func main() {
	te := editor.New(editor.NewTextWindow())
	te.Write("Salom Gopher!\n")
	te.Write("Ishlar qalay!")
	te.Save()

	fmt.Printf("Saqlangan matn:\n%v", te.Print())

	te.Write("Bu matn o'chirilib yuboriladi.")
	te.Undo()

	fmt.Printf("\nTiklangan matn:\n%v", te.Print())
}
