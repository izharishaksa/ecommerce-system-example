package use_case

import "github.com/google/uuid"

type ProductDetail struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	SalePrice float64   `json:"sale_price"`
	Stock     int       `json:"stock"`
	Sold      int       `json:"sold"`
}
