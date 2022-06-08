package use_case

import (
	"github.com/google/uuid"
	"inventory-service/internal/inventory"
)

type ProductDetail struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	SalePrice float64   `json:"sale_price"`
	Stock     int       `json:"stock"`
	Sold      int       `json:"sold"`
}

func fromProductToProductDetail(product inventory.Product) ProductDetail {
	return ProductDetail{
		ID:        product.Id,
		Title:     product.Title,
		SalePrice: product.SalePrice,
		Stock:     product.Stock,
		Sold:      product.Sold,
	}
}

type OrderAcceptedResponse struct {
	Id         uuid.UUID `json:"id"`
	CustomerId uuid.UUID `json:"customer_id"`
	TotalPrice float64   `json:"total_price"`
}

type OrderRejectedResponse struct {
	Id         uuid.UUID `json:"id"`
	CustomerId uuid.UUID `json:"customer_id"`
	Message    string    `json:"message"`
}
