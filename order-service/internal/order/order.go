package order

import (
	"github.com/google/uuid"
	"time"
)

type Status string

type Order struct {
	Id         uuid.UUID
	CustomerId uuid.UUID
	Items      []Item
	Status     Status
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

func PlaceOrder(customerId uuid.UUID, items []Item) (*Order, *OrderPlaced, error) {
	order := &Order{
		Id:         uuid.New(),
		CustomerId: customerId,
		Items:      items,
		Status:     Placed,
		TotalPrice: 0,
		CreatedAt:  time.Now(),
	}
	orderPlacedEvent := &OrderPlaced{
		EventType:  PlacedEvent,
		Id:         order.Id,
		CustomerId: order.CustomerId,
		Items: func(items []Item) []OrderItem {
			var orderItems []OrderItem
			for _, item := range items {
				orderItems = append(orderItems, OrderItem{
					ProductId: item.ProductId,
					Quantity:  item.Quantity,
				})
			}
			return orderItems
		}(order.Items),
		Status:     string(order.Status),
		TotalPrice: order.TotalPrice,
		CreatedAt:  order.CreatedAt,
		Message:    order.Message,
	}
	return order, orderPlacedEvent, nil
}
