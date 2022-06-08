package use_case

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"order-service/internal/inventory"
	"order-service/internal/order"
)

const (
	OrderStatusPlacedTopic   = "ORDER_PLACED"
	OrderStatusCreatedTopic  = "ORDER_CREATED"
	OrderStatusPaidTopic     = "ORDER_PAID"
	OrderStatusCanceledTopic = "ORDER_CANCELED"
	OrderStatusRejectedTopic = "ORDER_REJECTED"
)

type OrderService interface {
	CreateOrder(request CreateOrderRequest) (*uuid.UUID, error)
}

type orderService struct {
	orderRepository     order.Repository
	inventoryRepository inventory.Repository
	kafkaWriter         *kafka.Writer
}

func NewOrderService(orderRepository order.Repository, inventoryRepository inventory.Repository, kafkaWriter *kafka.Writer) OrderService {
	return &orderService{
		orderRepository:     orderRepository,
		inventoryRepository: inventoryRepository,
		kafkaWriter:         kafkaWriter,
	}
}

func (service orderService) CreateOrder(request CreateOrderRequest) (*uuid.UUID, error) {
	productItemIds := make([]uuid.UUID, len(request.Items))
	items := make([]order.Item, len(request.Items))
	for i, item := range request.Items {
		productItemIds[i] = item.ProductId
		items[i] = item.toOrderItem()
	}
	productDetails, err := service.inventoryRepository.GetProductAvailability(productItemIds)
	if err != nil {
		return nil, err
	}
	placedOrder, err := order.PlaceOrder(request.CustomerId, items, productDetails)
	if err != nil {
		return nil, err
	}
	err = service.orderRepository.SaveOrder(placedOrder)
	if err != nil {
		return nil, err
	}
	message := kafka.Message{
		Key:   []byte(placedOrder.Id.String()),
		Value: []byte(placedOrder.Status),
		Topic: OrderStatusPlacedTopic,
	}
	err = service.kafkaWriter.WriteMessages(context.Background(), message)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Message sent to topic %s: %s=%s\n", message.Topic, message.Key, message.Value)
	}
	return &placedOrder.Id, nil
}
