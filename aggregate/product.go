package aggregate

import (
	"errors"
	"github.com/furkandemireleng/go-ddd/entity"
)

type Product struct {
	item        *entity.Item
	name        string
	description string
	price       float64
	quantity    int
}

var (
	ErrMissingName        = errors.New("missing name")
	ErrMissingDescription = errors.New("missing description")
	ErrMissingPrice       = errors.New("missing price")
	ErrMissingQuantity    = errors.New("missing quantity")
)

// New Product Factory

func NewProduct(name string, description string, price float64, quantity int) (Product, error) {

	if name == "" {
		return Product{}, ErrMissingName
	}
	if description == "" {
		return Product{}, ErrMissingDescription
	}
	if price <= 0 {
		return Product{}, ErrMissingPrice
	}
	if quantity <= 0 {
		return Product{}, ErrMissingQuantity
	}

}
