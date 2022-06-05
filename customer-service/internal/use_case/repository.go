package use_case

import (
	"customer-service/internal/customer"
)

type CustomerRepository interface {
	SaveCustomer(customer *customer.Customer) error
}
