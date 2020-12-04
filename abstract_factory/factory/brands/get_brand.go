package brands

import (
	"fmt"

	"github.com/MrWebUzb/facade/abstract_factory/factory"
)

// GetSportsFactory ...
func GetSportsFactory(brand string) (factory.ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("unknown brand")
}
