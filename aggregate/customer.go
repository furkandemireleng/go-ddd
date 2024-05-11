// package aggregate holds our aggregates that combines many entities in to full object

package aggregate

import (
	"errors"
	"github.com/furkandemireleng/go-ddd/entity"
	"github.com/furkandemireleng/go-ddd/valueobject"
	"github.com/google/uuid"
)

var (
	ErrInvalidName = errors.New("Customer has to be a valid Name")
	ErrInvalidAge  = errors.New("Customer has to be a valid Age")
)

type Customer struct {
	// Person is the root entity of the customer
	// Which means person.id is the main identifier for the customer
	// I will make them lowercase because they are not accessible for other domains from outside

	// Aggregates should not be directly accessible from outside for grab data

	person   *entity.Person // I made pointer because  it could change states
	products []*entity.Item

	transactions []valueobject.Transaction // Transactions cant change so dont need pointer
}

// NewCustomer new customer is a Factory Pattern to create a new customer aggregate
// It will validate that name is not empty and age is not zero

func NewCustomer(name string, age int) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidName
	}
	if age <= 0 {
		return Customer{}, ErrInvalidAge
	}
	// let's create person entity
	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
		Age:  age,
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil

}
