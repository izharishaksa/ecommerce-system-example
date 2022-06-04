package use_case

import (
	"github.com/google/uuid"
	"inventory-service/internal/inventory"
)

type ProductService struct {
	productRepository ProductRepository
}

func NewProductService(productRepository ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (p *ProductService) CreateProduct(request CreateProductRequest) (*uuid.UUID, error) {
	product, err := inventory.NewProduct(request.Name, request.Price, request.Quantity)
	if err != nil {
		return nil, err
	}
	err = p.productRepository.SaveProduct(product)
	if err != nil {
		return nil, err
	}
	return &product.Id, nil
}

func (p *ProductService) GetAllProducts() ([]ProductDetail, error) {
	products, err := p.productRepository.GetAllProducts()
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
