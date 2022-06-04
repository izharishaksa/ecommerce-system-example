package use_case

import (
	"github.com/google/uuid"
	"inventory-service/internal/inventory"
)

type InventoryRepository interface {
	SaveProduct(product *inventory.Product) error
	FindProductById(id uuid.UUID) (*inventory.Product, error)
	GetAllProducts() ([]*inventory.Product, error)
}
