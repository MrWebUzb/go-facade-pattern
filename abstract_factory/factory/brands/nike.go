package brands

import "github.com/MrWebUzb/facade/abstract_factory/factory"

// Nike ...
type Nike struct{}

// NikeShoe ...
type NikeShoe struct{ factory.Shoe }

// NikeShirt ...
type NikeShirt struct{ factory.Shirt }

// MakeShoe ...
func (a *Nike) MakeShoe() factory.IShoe {
	s := &NikeShoe{}
	s.SetLogo("nike")
	s.SetSize(43)

	return s
}

// MakeShirt ...
func (a *Nike) MakeShirt() factory.IShirt {
	s := &NikeShirt{}
	s.SetLogo("nike")
	s.SetSize(38)

	return s
}
