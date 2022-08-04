package use_case

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"order-service/internal/order"
)

type OrderRepository interface {
	SaveOrder(order *order.Order) error
	GetAllOrders() ([]order.Order, error)
	FindOrderById(id uuid.UUID) (*order.Order, error)
}

type EventPublisher interface {
	Publish(ctx context.Context, topic string, id string, value []byte) error
}

type orderServiceImpl struct {
	orderRepository OrderRepository
	eventPublisher  EventPublisher
}

func NewOrderService(orderRepository OrderRepository, eventPublisher EventPublisher) *orderServiceImpl {
	return &orderServiceImpl{
		orderRepository: orderRepository,
		eventPublisher:  eventPublisher,
	}
}

func (service orderServiceImpl) PlaceOrder(request PlaceOrderRequest) (*uuid.UUID, error) {
	productItemIds := make([]uuid.UUID, len(request.Items))
	items := make([]order.Item, len(request.Items))
	for i, item := range request.Items {
		productItemIds[i] = item.ProductId
		items[i] = item.toOrderItem()
	}
	placedOrder, event, err := order.PlaceOrder(request.CustomerId, items)
	if err != nil {
		return nil, err
	}
	err = service.orderRepository.SaveOrder(placedOrder)
	if err != nil {
		return nil, err
	}
	serializedEvent, err := json.Marshal(event)
	err = service.eventPublisher.Publish(context.Background(), "Order", placedOrder.Id.String(), serializedEvent)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Event published into topic order: %s=%s\n", event.Id, serializedEvent)
	}
	return &placedOrder.Id, nil
}

func (service orderServiceImpl) OrderRejected(request order.OrderRejectedMessage) error {
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

func (service orderServiceImpl) OrderCreated(request order.OrderPlacedMessage) error {
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
