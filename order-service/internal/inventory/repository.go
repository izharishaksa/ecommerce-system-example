package inventory

import "github.com/google/uuid"

type Repository interface {
	GetProductAvailability(productIds []uuid.UUID) ([]ProductDetail, error)
}
