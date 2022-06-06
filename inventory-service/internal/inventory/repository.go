package inventory

import (
	"github.com/google/uuid"
)

type Repository interface {
	SaveProduct(product *Product) error
	FindProductById(id uuid.UUID) (*Product, error)
	GetAllProducts() ([]Product, error)
}
