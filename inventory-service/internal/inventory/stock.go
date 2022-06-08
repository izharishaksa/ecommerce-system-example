package inventory

import (
	"github.com/google/uuid"
	"lib"
)

type Stock struct {
	products map[uuid.UUID]*Product
}

func NewStock() *Stock {
	return &Stock{
		products: make(map[uuid.UUID]*Product),
	}
}

func (s Stock) UpdateByOrder(orderItems []OrderItem) (*float64, error) {
	totalPrice := 0.0
	for _, item := range orderItems {
		product, ok := s.products[item.ProductId]
		if !ok {
			return nil, lib.NewErrNotFound("product not found")
		}

		if err := product.DecreaseStock(item.Quantity); err != nil {
			return nil, err
		}
		totalPrice += product.SalePrice * float64(item.Quantity)
	}
	return &totalPrice, nil
}
