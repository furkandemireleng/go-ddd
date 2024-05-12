package services

import (
	"fmt"
	"github.com/furkandemireleng/go-ddd/aggregate"
	"github.com/furkandemireleng/go-ddd/domain/customer"
	"github.com/furkandemireleng/go-ddd/domain/customer/memory"
	"github.com/furkandemireleng/go-ddd/domain/product"
	prodmemory "github.com/furkandemireleng/go-ddd/domain/product/memory"
	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepositoty
	products  product.ProductRepository
}

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

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)

}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmemory.New()

		for _, product := range products {
			if err := pr.Add(product); err != nil {
				return nil
			}
		}

		os.products = pr
		return nil
	}

}

func (o *OrderService) CreateOrder(customerId uuid.UUID, productIds []uuid.UUID) error {
	// fetch the customer
	c, err := o.customers.Get(customerId)
	if err != nil {
		return err
	}
	fmt.Println(c)
	//get each product

	var productsArray []aggregate.Product
	var totalPrice float64

	for _, id := range productIds {
		p, err := o.products.GetByID(id)

		if err != nil {
			return err
		}

		productsArray = append(productsArray, p)
		totalPrice += p.GetPrice()
	}

	return nil
}
