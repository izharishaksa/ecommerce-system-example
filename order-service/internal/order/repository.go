package order

import "github.com/google/uuid"

type Repository interface {
	SaveOrder(order *Order) error
	GetAllOrders() ([]Order, error)
	FindOrderById(id uuid.UUID) (*Order, error)
}
