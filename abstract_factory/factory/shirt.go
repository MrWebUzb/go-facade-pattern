package factory

// IShirt ...
type IShirt interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

// Shirt ...
type Shirt struct {
	logo string
	size int
}

// SetLogo ...
func (s *Shirt) SetLogo(logo string) {
	s.logo = logo
}

// GetLogo ...
func (s Shirt) GetLogo() string {
	return s.logo
}

// SetSize ...
func (s *Shirt) SetSize(size int) {
	s.size = size
}

// GetSize ...
func (s Shirt) GetSize() int {
	return s.size
}
