package use_case

import "inventory-service/internal/inventory"

type ProductService struct {
	productRepository ProductRepository
}

func NewProductService(productRepository ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (p *ProductService) CreateProduct(request CreateProductRequest) error {
	product, err := inventory.NewProduct(request.Name, request.Price, request.Quantity)
	if err != nil {
		return err
	}
	return p.productRepository.SaveProduct(product)
}
