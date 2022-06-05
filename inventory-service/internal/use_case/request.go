package use_case

import "github.com/google/uuid"

type CreateProductRequest struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type AddStockRequest struct {
	ProductId uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	AtPrice   float64   `json:"at_price"`
}
