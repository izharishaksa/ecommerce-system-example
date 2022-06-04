package inventory

import (
	"github.com/google/uuid"
	"lib"
)

type InMemoryRepository struct {
	products map[uuid.UUID]*Product
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		products: make(map[uuid.UUID]*Product),
	}
}

func (repo *InMemoryRepository) SaveProduct(product *Product) error {
	repo.products[product.Id] = product
	return nil
}

func (repo *InMemoryRepository) FindProductById(id uuid.UUID) (*Product, error) {
	product, ok := repo.products[id]
	if !ok {
		return nil, lib.NewErrNotFound("product not found")
	}
	return product, nil
}

func (repo *InMemoryRepository) GetAllProducts() ([]*Product, error) {
	var products []*Product
	for _, product := range repo.products {
		products = append(products, product)
	}
	return products, nil
}
