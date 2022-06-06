package use_case

import (
	"github.com/google/uuid"
	"order-service/internal/order"
)

type OrderService struct {
	orderRepository     OrderRepository
	inventoryRepository InventoryRepository
}

func NewOrderService(orderRepository OrderRepository, inventoryRepository InventoryRepository) *OrderService {
	return &OrderService{
		orderRepository:     orderRepository,
		inventoryRepository: inventoryRepository,
	}
}

func (service OrderService) CreateOrder(request CreateOrderRequest) (*uuid.UUID, error) {
	productItemIds := make([]uuid.UUID, len(request.Items))
	items := make([]order.Item, len(request.Items))
	for i, item := range request.Items {
		productItemIds[i] = item.ProductId
		items[i] = order.Item{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		}
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
