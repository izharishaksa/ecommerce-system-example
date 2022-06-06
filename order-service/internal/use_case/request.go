package use_case

import (
	"github.com/google/uuid"
)

type CreateOrderRequest struct {
	CustomerId uuid.UUID   `json:"customer_id"`
	Items      []OrderItem `json:"items"`
}

type OrderItem struct {
	ProductId uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}
