package fs

// INode node interface of filesystem
type INode interface {
	Print(string)
	Clone() INode
}
