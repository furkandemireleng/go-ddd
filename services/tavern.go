package services

import "github.com/google/uuid"

type TavernCongiguration func(os *Tavern) error

type Tavern struct {

	// Tavern is a service which holds sub services

	OrderService *OrderService

	//for example like billing service
}

func NewTavernService(cfgs ...TavernCongiguration) (*Tavern, error) {
	t := &Tavern{}

	for _, cfg := range cfgs {
		if err := cfg(t); err != nil {
			return nil, err
		}
	}

	return t, nil
}

func WithOrderService(os *OrderService) TavernCongiguration {
	return func(t *Tavern) error {

		t.OrderService = os
		return nil
	}

}

func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {

	err := t.OrderService.CreateOrder(customer, products)

	if err != nil {
		return err
	}

	return nil
}
