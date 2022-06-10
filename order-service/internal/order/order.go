package order

import (
	"github.com/google/uuid"
	"time"
)

const (
	Placed   = "placed"
	Created  = "created"
	Rejected = "rejected"
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
	o.Status = Rejected
	o.Message = message
	return nil
}

func (o *Order) CreatedAtPrice(price float64) error {
	o.Status = Created
	o.TotalPrice = price
	return nil
}

func PlaceOrder(customerId uuid.UUID, items []Item) (*Order, error) {
	order := &Order{
		Id:         uuid.New(),
		CustomerId: customerId,
		Items:      items,
		Status:     Placed,
		TotalPrice: 0,
		CreatedAt:  time.Now(),
	}
	return order, nil
}
