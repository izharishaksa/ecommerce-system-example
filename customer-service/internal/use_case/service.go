package use_case

import (
	"customer-service/internal/customer"
	"github.com/google/uuid"
)

type CustomerService interface {
	RegisterCustomer(name, email string) (*uuid.UUID, error)
	GetAllCustomers() ([]CustomerDetail, error)
	TopUp(customerId uuid.UUID, amount float64) error
}

type customerService struct {
	customerRepository customer.Repository
}

func NewCustomerService(customerRepository customer.Repository) CustomerService {
	return &customerService{customerRepository: customerRepository}
}

func (service *customerService) RegisterCustomer(name, email string) (*uuid.UUID, error) {
	customerInstance, err := customer.NewCustomer(name, email)
	if err != nil {
		return nil, err
	}
	err = service.customerRepository.SaveCustomer(customerInstance)
	if err != nil {
		return nil, err
	}
	return &customerInstance.Id, err
}

func (service *customerService) GetAllCustomers() ([]CustomerDetail, error) {
	customers, err := service.customerRepository.GetCustomer()
	if err != nil {
		return nil, err
	}
	customerDetails := make([]CustomerDetail, 0)
	for _, cust := range customers {
		customerDetails = append(customerDetails, fromCustomerToCustomerDetail(cust))
	}
	return customerDetails, nil
}

func (service *customerService) TopUp(customerId uuid.UUID, amount float64) error {
	customerInstance, err := service.customerRepository.FindCustomerById(customerId)
	if err != nil {
		return err
	}
	err = customerInstance.TopUp(amount)
	if err != nil {
		return err
	}
	return service.customerRepository.UpdateBalance(customerInstance)
}
