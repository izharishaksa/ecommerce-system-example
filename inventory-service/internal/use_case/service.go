package use_case

import (
	"github.com/google/uuid"
	"inventory-service/internal/inventory"
)

type InventoryService struct {
	inventoryRepository InventoryRepository
}

func NewInventoryService(inventoryRepository InventoryRepository) *InventoryService {
	return &InventoryService{inventoryRepository: inventoryRepository}
}

func (service *InventoryService) CreateProduct(request CreateProductRequest) (*uuid.UUID, error) {
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

func (service *InventoryService) GetAllProducts() ([]ProductDetail, error) {
	products, err := service.inventoryRepository.GetAllProducts()
	if err != nil {
		return nil, err
	}
	var productDetails []ProductDetail
	for _, product := range products {
		productDetails = append(productDetails, ProductDetail{
			ID:        product.Id,
			Title:     product.Title,
			SalePrice: product.SalePrice,
			Stock:     product.Stock,
			Sold:      product.Sold,
		})
	}
	return productDetails, nil
}

func (service *InventoryService) AddStock(request AddStockRequest) error {
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
