package use_case

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"inventory-service/internal/inventory"
)

const (
	OrderRejected = "ORDER_REJECTED"
	OrderCreated  = "ORDER_CREATED"
)

type InventoryRepository interface {
	SaveProduct(product *inventory.Product) error
	FindProductById(id uuid.UUID) (*inventory.Product, error)
	GetAllProducts() ([]inventory.Product, error)
	GetAvailableStock(ids []uuid.UUID) (*inventory.Stock, error)
	UpdateStock(stock *inventory.Stock) error
}

type EventPublisher interface {
	Publish(ctx context.Context, event string, id string, value []byte) error
}

type inventoryService struct {
	inventoryRepository InventoryRepository
	eventPublisher      EventPublisher
}

func NewInventoryService(inventoryRepository InventoryRepository, eventPublisher EventPublisher) *inventoryService {
	return &inventoryService{inventoryRepository: inventoryRepository, eventPublisher: eventPublisher}
}

func (service *inventoryService) CreateProduct(request CreateProductRequest) (*uuid.UUID, error) {
	fmt.Println()
	product, err := inventory.NewProduct(request.Title, request.Price, request.Quantity)
	if err != nil {
		return nil, err
	}
	err = service.inventoryRepository.SaveProduct(product)
	if err != nil {
		return nil, err
	}
	return &product.Id, nil
}

func (service *inventoryService) GetAllProducts() ([]ProductDetail, error) {
	products, err := service.inventoryRepository.GetAllProducts()
	if err != nil {
		return nil, err
	}
	var productDetails []ProductDetail
	for _, product := range products {
		productDetails = append(productDetails, fromProductToProductDetail(product))
	}
	return productDetails, nil
}

func (service *inventoryService) AddStock(request AddStockRequest) error {
	product, err := service.inventoryRepository.FindProductById(request.ProductId)
	if err != nil {
		return err
	}
	err = product.AddStock(request.Quantity, request.AtPrice)
	if err != nil {
		return err
	}
	return service.inventoryRepository.SaveProduct(product)
}

func (service *inventoryService) OrderPlaced(request PlacedOrderRequest) error {
	orderItems := make([]inventory.OrderItem, 0)
	productIds := make([]uuid.UUID, 0)
	for _, item := range request.Items {
		orderItems = append(orderItems, item.toOrderItem())
		productIds = append(productIds, item.ProductId)
	}
	var err error
	var totalPrice *float64
	stock, err := service.inventoryRepository.GetAvailableStock(productIds)
	rejectedResponse := OrderRejectedResponse{
		Id:         request.Id,
		CustomerId: request.CustomerId,
	}
	if err != nil {
		rejectedResponse.Message = err.Error()
		messageValue, err := json.Marshal(rejectedResponse)
		if err != nil {
			return err
		}
		return service.eventPublisher.Publish(context.Background(), OrderRejected, request.Id.String(), messageValue)
	}
	totalPrice, err = stock.UpdateByOrder(orderItems)
	if err != nil {
		rejectedResponse.Message = err.Error()
		messageValue, err := json.Marshal(rejectedResponse)
		if err != nil {
			return err
		}
		return service.eventPublisher.Publish(context.Background(), OrderRejected, request.Id.String(), messageValue)
	}
	response := OrderAcceptedResponse{
		Id:         request.Id,
		CustomerId: request.CustomerId,
		TotalPrice: *totalPrice,
	}
	messageValue, err := json.Marshal(response)
	if err != nil {
		return err
	}
	err = service.inventoryRepository.UpdateStock(stock)
	if err != nil {
		return err
	}
	return service.eventPublisher.Publish(context.Background(), OrderCreated, request.Id.String(), messageValue)
}
