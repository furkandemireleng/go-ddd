package services

import (
	"github.com/furkandemireleng/go-ddd/aggregate"
	"github.com/google/uuid"
	"testing"
)

func init_products(t *testing.T) []aggregate.Product {

	beer, err := aggregate.NewProduct("Beer", "beer", 10, 2)

	if err != nil {
		t.Fatal(err)
	}

	peenut, err := aggregate.NewProduct("People", "people", 13.99, 2)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := aggregate.NewProduct("Wine", "wine", 12.99, 2)
	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{beer, peenut, wine}
}

func TestOrder(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	customer, err := aggregate.NewCustomer("furkan", 12)
	if err != nil {
		t.Fatal(err)
	}

	customerErr := os.customers.Add(customer)
	if customerErr != nil {
		t.Fatal(err)
	}

	var productIds []uuid.UUID

	for _, product := range products {
		productIds = append(productIds, product.GetId())

	}

	orderErr := os.CreateOrder(customer.GetId(), productIds)

	if orderErr != nil {
		t.Fatal(orderErr)
	}

}
