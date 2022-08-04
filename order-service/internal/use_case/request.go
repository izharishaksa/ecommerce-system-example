package use_case

import (
	"github.com/google/uuid"
	"order-service/internal/order"
)

type PlaceOrderRequest struct {
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
