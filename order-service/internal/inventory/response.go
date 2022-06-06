package inventory

import "github.com/google/uuid"

type ProductDetail struct {
	ProductId    uuid.UUID
	CurrentPrice float64
	Stock        int
}
