package factory

// IShoe ...
type IShoe interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

// Shoe ...
type Shoe struct {
	logo string
	size int
}

// SetLogo ...
func (s *Shoe) SetLogo(logo string) {
	s.logo = logo
}

// GetLogo ...
func (s Shoe) GetLogo() string {
	return s.logo
}

// SetSize ...
func (s *Shoe) SetSize(size int) {
	s.size = size
}

// GetSize ...
func (s Shoe) GetSize() int {
	return s.size
}
