package order

import (
	"github.com/google/uuid"
	"lib"
)

type inMemoryRepository struct {
	orders map[string]*Order
}

func NewInMemoryRepository() Repository {
	return &inMemoryRepository{
		orders: make(map[string]*Order),
	}
}

func (repo inMemoryRepository) FindOrderById(id uuid.UUID) (*Order, error) {
	order, ok := repo.orders[id.String()]
	if !ok {
		return nil, lib.NewErrNotFound("order not found")
	}
	return order, nil
}

func (repo inMemoryRepository) SaveOrder(order *Order) error {
	repo.orders[order.Id.String()] = order
	return nil
}

func (repo inMemoryRepository) GetAllOrders() ([]Order, error) {
	var orders []Order
	for _, order := range repo.orders {
		orders = append(orders, *order)
	}
	return orders, nil
}
