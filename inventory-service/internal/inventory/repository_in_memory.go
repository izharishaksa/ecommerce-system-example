package inventory

import (
	"github.com/google/uuid"
	"lib"
)

type inMemoryRepository struct {
	products map[uuid.UUID]*Product
}

func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{
		products: make(map[uuid.UUID]*Product),
	}
}

func (repo *inMemoryRepository) SaveProduct(product *Product) error {
	repo.products[product.Id] = product
	return nil
}

func (repo *inMemoryRepository) FindProductById(id uuid.UUID) (*Product, error) {
	product, ok := repo.products[id]
	if !ok {
		return nil, lib.NewErrNotFound("product not found")
	}
	return product, nil
}

func (repo *inMemoryRepository) GetAllProducts() ([]Product, error) {
	var products []Product
	for _, product := range repo.products {
		products = append(products, *product)
	}
	return products, nil
}

func (repo *inMemoryRepository) GetAvailableStock(ids []uuid.UUID) (*Stock, error) {
	stock := NewStock()
	for _, id := range ids {
		product, ok := repo.products[id]
		if !ok {
			return nil, lib.NewErrNotFound("product not found")
		}

		stock.products[id] = product
	}
	return stock, nil
}

func (repo *inMemoryRepository) UpdateStock(stock *Stock) error {
	for id, product := range stock.products {
		oldProduct := repo.products[id]
		oldProduct.Stock = product.Stock
		oldProduct.Sold = product.Sold
		repo.products[id] = oldProduct
	}
	return nil
}
