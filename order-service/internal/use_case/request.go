package use_case

import (
	"github.com/google/uuid"
	"order-service/internal/order"
)

type CreateOrderRequest struct {
	CustomerId uuid.UUID          `json:"customer_id"`
	Items      []OrderItemRequest `json:"items"`
}

type OrderItemRequest struct {
	ProductId uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

func (r OrderItemRequest) toOrderItem() order.Item {
	return order.Item{
		ProductId: r.ProductId,
		Quantity:  r.Quantity,
	}
}

type OrderRejectedRequest struct {
	Id         uuid.UUID `json:"id"`
	CustomerId uuid.UUID `json:"customer_id"`
	Message    *string   `json:"message"`
}

type OrderCreatedRequest struct {
	Id         uuid.UUID `json:"id"`
	CustomerId uuid.UUID `json:"customer_id"`
	TotalPrice float64   `json:"total_price"`
}
