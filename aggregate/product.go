package aggregate

import (
	"errors"
	"github.com/furkandemireleng/go-ddd/entity"
	"github.com/google/uuid"
)

type Product struct {
	item     *entity.Item
	price    float64
	quantity int
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

	return Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: quantity,
	}, nil
}

func (p Product) GetId() uuid.UUID {
	return p.item.ID
}
func (p Product) GetName() string {
	return p.item.Name
}
func (p Product) GetPrice() float64 {
	return p.price
}
func (p Product) GetQuantity() int {
	return p.quantity
}
func (p Product) GetItem() *entity.Item {
	return p.item
}
