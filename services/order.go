package services

import (
	"fmt"
	"github.com/furkandemireleng/go-ddd/domain/customer"
	"github.com/google/uuid"
)

type OrderService struct {
	customers customer.CustomerRepositoty
}

type OrderConfiguration func(os *OrderService) error

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	//loop all the congigs and applied
	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

// with customer repository.go applies a customer repository.go to the order service

func WithCustomerRepository(cr customer.CustomerRepositoty) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

//func WithMemoryCustomerRepository() OrderConfiguration {
//	cr := memory.New()
//	return WithCustomerRepository(cr)
//
//}

func (o *OrderService) CreateOrder(customerId uuid.UUID, productIds []uuid.UUID) error {
	// fetch the customer
	c, err := o.customers.Get(customerId)
	if err != nil {
		return err
	}
	fmt.Println(c)

	return nil
}
