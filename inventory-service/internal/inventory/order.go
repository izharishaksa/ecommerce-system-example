package inventory

import "github.com/google/uuid"

type OrderItem struct {
	ProductId uuid.UUID
	Quantity  int
}
