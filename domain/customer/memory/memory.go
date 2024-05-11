package memory

import (
	"github.com/furkandemireleng/go-ddd/aggregate"
	"github.com/furkandemireleng/go-ddd/domain/customer"
	"github.com/google/uuid"
	"sync"
)

// package memory is a  n in-memory implementation of customer repository.go

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

// Factory pattern for adding new

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	if _, ok := mr.customers[c.GetId()]; !ok {
		return customer.ErrCustomerNotFound
	}

	// update the customer

	mr.Lock()
	mr.customers[c.GetId()] = c
	mr.Unlock()

	return nil
}

func (mr *MemoryRepository) Delete(id uuid.UUID) error {
	mr.Lock()

	if _, ok := mr.customers[id]; !ok {
		return customer.ErrCustomerNotFound
	}

	delete(mr.customers, id)

	mr.Unlock()

	return nil
}
