package memory

import (
	"github.com/furkandemireleng/go-ddd/aggregate"
	product2 "github.com/furkandemireleng/go-ddd/domain/product"
	"github.com/google/uuid"
	"sync"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mpr *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product

	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
}

func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	product, ok := mpr.products[id]
	if !ok {
		return aggregate.Product{}, product2.ErrProductNotFound
	}
	return product, nil
}

func (mpr *MemoryProductRepository) Add(p aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[p.GetId()]; ok {
		return product2.ErrProductAlreadyExists
	}
	mpr.products[p.GetId()] = p

	return nil

}

func (mpr *MemoryProductRepository) Update(updatedProduct aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[updatedProduct.GetId()]; !ok {
		return product2.ErrProductNotFound
	}

	mpr.products[updatedProduct.GetId()] = updatedProduct
	return nil

}

func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product2.ErrProductNotFound
	}

	delete(mpr.products, id)

	return nil

}
