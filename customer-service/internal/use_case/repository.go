package use_case

import (
	"customer-service/internal/customer"
	"github.com/google/uuid"
)

type CustomerRepository interface {
	SaveCustomer(customer *customer.Customer) error
	GetCustomer() ([]customer.Customer, error)
	FindCustomerById(id uuid.UUID) (*customer.Customer, error)
}
