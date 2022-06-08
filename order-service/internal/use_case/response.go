package use_case

import (
	"github.com/google/uuid"
	"order-service/internal/order"
	"time"
)

type OrderResponse struct {
	Id         uuid.UUID      `json:"id"`
	CustomerId uuid.UUID      `json:"customer_id"`
	Items      []ItemResponse `json:"items"`
	Status     string         `json:"status"`
	TotalPrice float64        `json:"total_price"`
	CreatedAt  time.Time      `json:"created_at"`
	Message    *string        `json:"message"`
}

type ItemResponse struct {
	ProductId uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

func fromOrderToOrderDetail(order order.Order) OrderResponse {
	return OrderResponse{
		Id:         order.Id,
		CustomerId: order.CustomerId,
		Items:      fromOrderItemsToItemResponses(order.Items),
		Status:     order.Status,
		TotalPrice: order.TotalPrice,
		CreatedAt:  order.CreatedAt,
		Message:    order.Message,
	}
}

func fromOrderItemsToItemResponses(items []order.Item) []ItemResponse {
	var itemResponses []ItemResponse
	for _, item := range items {
		itemResponses = append(itemResponses, ItemResponse{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}
	return itemResponses
}
