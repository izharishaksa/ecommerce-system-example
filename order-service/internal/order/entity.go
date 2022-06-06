package order

import (
	"fmt"
	"github.com/google/uuid"
	"lib"
	"order-service/internal/inventory"
	"time"
)

const (
	OrderStatusPlaced   = "placed"
	OrderStatusCreated  = "created"
	OrderStatusPaid     = "paid"
	OrderStatusCanceled = "canceled"
)

type Order struct {
	Id         uuid.UUID
	CustomerId uuid.UUID
	Items      []Item
	Status     string
	TotalPrice float64
	CreatedAt  time.Time
}

type Item struct {
	ProductId uuid.UUID
	Quantity  int
}

func PlaceOrder(customerId uuid.UUID, items []Item, productDetail []inventory.ProductDetail) (*Order, error) {
	totalPrice := 0.0
	for _, item := range items {
		isAvailable := false
		for _, product := range productDetail {
			if item.ProductId == product.ProductId {
				isAvailable = true
				if product.Stock < item.Quantity {
					return nil, lib.NewErrBadRequest("Not enough stock")
				}
				totalPrice += product.CurrentPrice * float64(item.Quantity)
			}
		}
		if !isAvailable {
			return nil, lib.NewErrBadRequest(fmt.Sprintf("Product %s is not found", item.ProductId))
		}
	}
	order := &Order{
		Id:         uuid.New(),
		CustomerId: customerId,
		Items:      items,
		Status:     OrderStatusPlaced,
		TotalPrice: totalPrice,
		CreatedAt:  time.Now(),
	}
	return order, nil
}
