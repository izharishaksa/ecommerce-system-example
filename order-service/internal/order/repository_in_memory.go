package order

import (
	"github.com/google/uuid"
	"lib"
	"sync"
)

type inMemoryRepository struct {
	mu     sync.RWMutex
	orders map[string]*Order
}

func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{
		orders: make(map[string]*Order),
	}
}

func (repo *inMemoryRepository) FindOrderById(id uuid.UUID) (*Order, error) {
	repo.mu.RLock()
	order, ok := repo.orders[id.String()]
	repo.mu.RUnlock()
	if !ok {
		return nil, lib.NewErrNotFound("order not found")
	}
	return order, nil
}

func (repo *inMemoryRepository) SaveOrder(order *Order) error {
	repo.mu.Lock()
	repo.orders[order.Id.String()] = order
	repo.mu.Unlock()
	return nil
}

func (repo *inMemoryRepository) GetAllOrders() ([]Order, error) {
	repo.mu.RLock()
	var orders []Order
	for _, order := range repo.orders {
		orders = append(orders, *order)
	}
	repo.mu.RUnlock()
	return orders, nil
}
