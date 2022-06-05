package use_case

import (
	"customer-service/internal/customer"
	"github.com/google/uuid"
)

type CustomerService struct {
	customerRepository CustomerRepository
}

func NewCustomerService(customerRepository CustomerRepository) *CustomerService {
	return &CustomerService{customerRepository: customerRepository}
}

func (service *CustomerService) RegisterCustomer(name string) (*uuid.UUID, error) {
	customerInstance, err := customer.NewCustomer(name)
	if err != nil {
		return nil, err
	}
	err = service.customerRepository.SaveCustomer(customerInstance)
	if err != nil {
		return nil, err
	}
	return &customerInstance.Id, err
}
