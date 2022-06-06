package customer

import (
	"github.com/google/uuid"
)

type Repository interface {
	SaveCustomer(customer *Customer) error
	GetCustomer() ([]Customer, error)
	FindCustomerById(id uuid.UUID) (*Customer, error)
	UpdateBalance(customer *Customer) error
}
