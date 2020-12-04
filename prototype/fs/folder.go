package fs

import "fmt"

// Folder ...
type Folder struct {
	Childs []INode
	Name   string
}

// Print ...
func (f *Folder) Print(indent string) {
	fmt.Printf("%s%s\n", indent, f.Name)
	for _, file := range f.Childs {
		file.Print(indent + indent)
	}
}

// Clone ...
func (f *Folder) Clone() INode {
	clone := &Folder{Name: f.Name + "_clone"}
	var tmpChilds []INode

	for _, file := range f.Childs {
		copy := file.Clone()
		tmpChilds = append(tmpChilds, copy)
	}
	clone.Childs = tmpChilds
	return clone
}
