package fs

import "fmt"

// File ...
type File struct {
	Name string
}

// Print ...
func (f *File) Print(indent string) {
	fmt.Printf("%s%s_name\n", indent, f.Name)
}

// Clone ...
func (f *File) Clone() INode {
	return &File{Name: fmt.Sprintf("%s_clone", f.Name)}
}
