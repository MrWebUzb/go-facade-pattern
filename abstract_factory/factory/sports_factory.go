package factory

// ISportsFactory ...
type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}
