package event

import (
	"github.com/google/uuid"
	"time"
)

type Type string

const (
	OrderPlacedType Type = "order_placed"
)

type OrderPlaced struct {
	EventType  string      `json:"event_type"`
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
	EventType  string    `json:"event_type"`
	Id         uuid.UUID `json:"id"`
	CustomerId uuid.UUID `json:"customer_id"`
	Message    *string   `json:"message"`
}

type OrderPlacedMessage struct {
	EventType  string    `json:"event_type"`
	Id         uuid.UUID `json:"id"`
	CustomerId uuid.UUID `json:"customer_id"`
	TotalPrice float64   `json:"total_price"`
}
