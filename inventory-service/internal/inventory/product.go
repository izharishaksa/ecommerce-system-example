package inventory

import (
	"github.com/google/uuid"
	"lib"
	"strings"
)

type Product struct {
	Id           uuid.UUID
	Title        string
	SalePrice    float64
	AveragePrice float64
	Stock        int
	Sold         int
}

func NewProduct(title string, price float64, quantity int) (*Product, error) {
	instance := &Product{
		Id:           uuid.New(),
		SalePrice:    0,
		AveragePrice: 0,
		Stock:        0,
		Sold:         0,
	}
	if err := instance.SetTitle(title); err != nil {
		return nil, err
	}
	if err := instance.AddStock(quantity, price); err != nil {
		return nil, err
	}
	return instance, nil
}

func (p *Product) SetTitle(newTitle string) error {
	if strings.TrimSpace(newTitle) == "" {
		return lib.NewErrBadRequest("title cannot be empty")
	}
	p.Title = newTitle
	return nil
}

func (p *Product) AddStock(quantity int, price float64) error {
	if quantity < 1 {
		return lib.NewErrBadRequest("quantity must be greater than 0")
	}
	if price <= 0 {
		return lib.NewErrBadRequest("price must be greater than 0")
	}
	currentTotalPrice := float64(p.Stock) * p.AveragePrice
	p.Stock += quantity
	p.AveragePrice = (currentTotalPrice + float64(quantity)*price) / float64(p.Stock)

	if p.SalePrice < p.AveragePrice {
		p.SalePrice = p.AveragePrice
	}
	return nil
}

func (p *Product) UpdateSalePrice(newSalePrice float64) error {
	if newSalePrice < p.AveragePrice {
		return lib.NewErrBadRequest("new sale price cannot be less than average price")
	}
	p.SalePrice = newSalePrice
	return nil
}

func (p *Product) DecreaseStock(quantity int) error {
	if p.Stock < quantity {
		return lib.NewErrBadRequest("not enough stock")
	}
	p.Stock -= quantity
	p.Sold += quantity
	return nil
}
