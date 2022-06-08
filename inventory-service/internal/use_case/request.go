package use_case

import (
	"github.com/google/uuid"
	"inventory-service/internal/inventory"
)

type CreateProductRequest struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type AddStockRequest struct {
	ProductId uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	AtPrice   float64   `json:"at_price"`
}

type PlacedOrderRequest struct {
	Id         uuid.UUID          `json:"id"`
	CustomerId uuid.UUID          `json:"customer_id"`
	Items      []OrderItemRequest `json:"items"`
}

type OrderItemRequest struct {
	ProductId uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

func (o OrderItemRequest) toOrderItem() inventory.OrderItem {
	return inventory.OrderItem{
		ProductId: o.ProductId,
		Quantity:  o.Quantity,
	}
}
