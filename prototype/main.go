package main

import (
	"fmt"

	"github.com/MrWebUzb/facade/prototype/fs"
)

func main() {
	music1 := &fs.File{Name: "Abdulla Qurbonov.mp3"}
	music2 := &fs.File{Name: "Kichkintoy qiz.mp3"}
	music3 := &fs.File{Name: "Andijonlik.mp3"}
	doc1 := &fs.File{Name: "DTTL Mustaqil ish.pdf"}

	music := &fs.Folder{
		Name:   "Music",
		Childs: []fs.INode{music1, music2, music3},
	}

	documents := &fs.Folder{
		Name:   "Documents",
		Childs: []fs.INode{music, doc1},
	}

	fmt.Println("Print hierarchy of Documents")
	documents.Print("--")

	// clone Documents folder
	cloneFolder := documents.Clone()
	fmt.Println("\nPrinting hierarchy for cloned Documents")
	cloneFolder.Print("--")

}
