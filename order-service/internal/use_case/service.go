package use_case

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"order-service/internal/order"
)

const (
	OrderStatusPlacedTopic   = "ORDER_PLACED"
	OrderStatusCreatedTopic  = "ORDER_CREATED"
	OrderStatusPaidTopic     = "ORDER_PAID"
	OrderStatusCanceledTopic = "ORDER_CANCELED"
	OrderStatusRejectedTopic = "ORDER_REJECTED"
)

type repository interface {
	SaveOrder(order *order.Order) error
	GetAllOrders() ([]order.Order, error)
	FindOrderById(id uuid.UUID) (*order.Order, error)
}

type orderServiceImpl struct {
	orderRepository repository
	kafkaWriter     *kafka.Writer
}

func NewOrderService(orderRepository repository, kafkaWriter *kafka.Writer) *orderServiceImpl {
	return &orderServiceImpl{
		orderRepository: orderRepository,
		kafkaWriter:     kafkaWriter,
	}
}

func (service orderServiceImpl) CreateOrder(request CreateOrderRequest) (*uuid.UUID, error) {
	productItemIds := make([]uuid.UUID, len(request.Items))
	items := make([]order.Item, len(request.Items))
	for i, item := range request.Items {
		productItemIds[i] = item.ProductId
		items[i] = item.toOrderItem()
	}
	placedOrder, err := order.PlaceOrder(request.CustomerId, items)
	if err != nil {
		return nil, err
	}
	err = service.orderRepository.SaveOrder(placedOrder)
	if err != nil {
		return nil, err
	}
	messageValue, err := json.Marshal(fromOrderToOrderDetail(*placedOrder))
	message := kafka.Message{
		Key:   []byte(placedOrder.Id.String()),
		Value: messageValue,
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

func (service orderServiceImpl) OrderRejected(request OrderRejectedRequest) error {
	orderById, err := service.orderRepository.FindOrderById(request.Id)
	if err != nil {
		return err
	}
	err = orderById.Reject(request.Message)
	if err != nil {
		return err
	}
	err = service.orderRepository.SaveOrder(orderById)
	if err != nil {
		return err
	}
	return nil
}

func (service orderServiceImpl) OrderCreated(request OrderCreatedRequest) error {
	orderById, err := service.orderRepository.FindOrderById(request.Id)
	if err != nil {
		return err
	}
	err = orderById.CreatedAtPrice(request.TotalPrice)
	if err != nil {
		return err
	}
	err = service.orderRepository.SaveOrder(orderById)
	if err != nil {
		return err
	}
	return nil
}

func (service orderServiceImpl) GetAllOrders() ([]OrderResponse, error) {
	orders, err := service.orderRepository.GetAllOrders()
	if err != nil {
		return nil, err
	}
	var orderDetails []OrderResponse
	for _, o := range orders {
		orderDetails = append(orderDetails, fromOrderToOrderDetail(o))
	}
	return orderDetails, nil
}
