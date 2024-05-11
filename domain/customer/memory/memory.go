package memory

import (
	"github.com/furkandemireleng/go-ddd/aggregate"
	"github.com/google/uuid"
	"sync"
)

// package memory is a  n in-memory implementation of customer repository

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
	return aggregate.Customer{}, nil
}

func (mr *MemoryRepository) Add(aggregate.Customer) error {
	return nil
}

func (mr *MemoryRepository) Update(aggregate.Customer) error {
	return nil
}

func (mr *MemoryRepository) Delete(uuid2 uuid.UUID) error {
	return nil
}
