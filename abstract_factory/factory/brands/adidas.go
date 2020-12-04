package brands

import "github.com/MrWebUzb/facade/abstract_factory/factory"

// Adidas ...
type Adidas struct{}

// AdidasShoe ...
type AdidasShoe struct{ factory.Shoe }

// AdidasShirt ...
type AdidasShirt struct{ factory.Shirt }

// MakeShoe ...
func (a *Adidas) MakeShoe() factory.IShoe {
	s := &AdidasShoe{Shoe: factory.Shoe{}}
	s.Shoe.SetLogo("adidas")
	s.Shoe.SetSize(41)

	return s
}

// MakeShirt ...
func (a *Adidas) MakeShirt() factory.IShirt {
	s := &AdidasShirt{Shirt: factory.Shirt{}}
	s.Shirt.SetLogo("adidas")
	s.Shirt.SetSize(40)

	return s
}
