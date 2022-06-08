package order

import (
	"github.com/google/uuid"
	"time"
)

const (
	OrderStatusPlaced   = "placed"
	OrderStatusCreated  = "created"
	OrderStatusRejected = "rejected"
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
	Message    *string
}

type Item struct {
	ProductId uuid.UUID
	Quantity  int
}

func (o *Order) Reject(message *string) error {
	o.Status = OrderStatusRejected
	o.Message = message
	return nil
}

func (o *Order) CreatedAtPrice(price float64) error {
	o.Status = OrderStatusCreated
	o.TotalPrice = price
	return nil
}

func PlaceOrder(customerId uuid.UUID, items []Item) (*Order, error) {
	order := &Order{
		Id:         uuid.New(),
		CustomerId: customerId,
		Items:      items,
		Status:     OrderStatusPlaced,
		TotalPrice: 0,
		CreatedAt:  time.Now(),
	}
	return order, nil
}
