package use_case

import (
	"customer-service/internal/customer"
	"github.com/google/uuid"
)

type CustomerDetail struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Balance float64   `json:"balance"`
}

func fromCustomerToCustomerDetail(customer customer.Customer) CustomerDetail {
	return CustomerDetail{
		ID:      customer.Id,
		Name:    customer.Name,
		Balance: customer.Balance,
	}
}
