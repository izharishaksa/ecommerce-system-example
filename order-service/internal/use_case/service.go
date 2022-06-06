package use_case

import (
	"github.com/google/uuid"
	"order-service/internal/inventory"
	"order-service/internal/order"
)

type OrderService interface {
	CreateOrder(request CreateOrderRequest) (*uuid.UUID, error)
}

type orderService struct {
	orderRepository     order.Repository
	inventoryRepository inventory.Repository
}

func NewOrderService(orderRepository order.Repository, inventoryRepository inventory.Repository) OrderService {
	return &orderService{
		orderRepository:     orderRepository,
		inventoryRepository: inventoryRepository,
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
	return &placedOrder.Id, nil
}
