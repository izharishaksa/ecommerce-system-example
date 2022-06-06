package use_case

import (
	"github.com/google/uuid"
	"order-service/internal/inventory"
	"order-service/internal/order"
)

type OrderRepository interface {
	SaveOrder(order *order.Order) error
}

type InventoryRepository interface {
	GetProductAvailability(productIds []uuid.UUID) ([]inventory.ProductDetail, error)
}
