package order

import (
	"github.com/google/uuid"
	"time"
)

type EventType string

type OrderPlaced struct {
	EventType  EventType   `json:"event_type"`
	Id         uuid.UUID   `json:"id"`
	CustomerId uuid.UUID   `json:"customer_id"`
	Items      []OrderItem `json:"items"`
	Status     string      `json:"status"`
	TotalPrice float64     `json:"total_price"`
	CreatedAt  time.Time   `json:"created_at"`
	Message    *string     `json:"message"`
}

type OrderItem struct {
	ProductId uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

type OrderRejectedMessage struct {
	EventType  EventType `json:"event_type"`
	Id         uuid.UUID `json:"id"`
	CustomerId uuid.UUID `json:"customer_id"`
	Message    *string   `json:"message"`
}

type OrderPlacedMessage struct {
	EventType  EventType `json:"event_type"`
	Id         uuid.UUID `json:"id"`
	CustomerId uuid.UUID `json:"customer_id"`
	TotalPrice float64   `json:"total_price"`
}
