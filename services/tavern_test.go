package services

import (
	"github.com/furkandemireleng/go-ddd/aggregate"
	"github.com/google/uuid"
	"testing"
)

func TestTavern(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	tavernService, tavernErr := NewTavernService(
		WithOrderService(os))

	if tavernErr != nil {
		t.Fatal(tavernErr)
	}

	customer, customerErr := aggregate.NewCustomer("Furkan", 12)

	if customerErr != nil {
		t.Fatal(customerErr)
	}

	var productsIds []uuid.UUID
	for _, product := range products {
		productsIds = append(productsIds, product.GetId())

	}
	completeOrder := tavernService.Order(customer.GetId(), productsIds)

	if completeOrder == nil {
		t.Fatal("order should not be nil")
	}
}
