package customer

import (
	"errors"
	"github.com/furkandemireleng/go-ddd/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound       = errors.New("Customer not found")
	ErrCustomerAlreadyExists  = errors.New("Customer already exists")
	ErrFailedToCreateCustomer = errors.New("Failed to create customer")
	ErrFailedToUpdateCustomer = errors.New("Failed to update customer")
	ErrFailedToDeleteCustomer = errors.New("Failed to delete customer")
)

type CustomerRepositoty interface {
	Get(uuid uuid.UUID) (aggregate.Customer, error)
	Add(customer aggregate.Customer) error
	Update(customer aggregate.Customer) error
	Delete(uuid uuid.UUID) error
}
