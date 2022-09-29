package use_case

import (
	"github.com/google/uuid"
	"order-service/internal/order"
	"time"
)

type OrderResponse struct {
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

func fromOrderToOrderDetail(order order.Order) OrderResponse {
	return OrderResponse{
		Id:         order.Id,
		CustomerId: order.CustomerId,
		Items:      fromOrderItemsToItemResponses(order.Items),
		Status:     string(order.Status),
		TotalPrice: order.TotalPrice,
		CreatedAt:  order.CreatedAt,
		Message:    order.Message,
	}
}

func fromOrderItemsToItemResponses(items []order.Item) []OrderItem {
	var orderItems []OrderItem
	for _, item := range items {
		orderItems = append(orderItems, OrderItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}
	return orderItems
}
