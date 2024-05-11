package product

import (
	"errors"
	"github.com/furkandemireleng/go-ddd/aggregate"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound      = errors.New("product not found")
	ErrProductAlreadyExists = errors.New("product already exists")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(p aggregate.Product) (aggregate.Product, error)
	Update(p aggregate.Product) (aggregate.Product, error)
	Delete(id uuid.UUID) error
}
