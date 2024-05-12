package memory

import (
	"fmt"
	"github.com/furkandemireleng/go-ddd/aggregate"
	"github.com/furkandemireleng/go-ddd/domain/customer"
	"github.com/google/uuid"
	"sync"
)

// package memory is an  in in-memory implementation of customer repository.go

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

// Add will add a new customer to the repository
func (mr *MemoryRepository) Add(c aggregate.Customer) error {
	if mr.customers == nil {
		// Saftey check if customers is not create, shouldn't happen if using the Factory, but you never know
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}
	// Make sure Customer isn't already in the repository
	if _, ok := mr.customers[c.GetId()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToCreateCustomer)
	}
	mr.Lock()
	mr.customers[c.GetId()] = c
	mr.Unlock()
	return nil
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
