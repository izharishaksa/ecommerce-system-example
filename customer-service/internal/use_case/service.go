package use_case

import (
	"customer-service/internal/customer"
	"github.com/google/uuid"
)

type customerRepository interface {
	SaveCustomer(customer *customer.Customer) error
	GetCustomer() ([]customer.Customer, error)
	FindCustomerById(id uuid.UUID) (*customer.Customer, error)
	UpdateBalance(customer *customer.Customer) error
}

type customerServiceImpl struct {
	customerRepository customerRepository
}

func NewCustomerService(customerRepository customerRepository) *customerServiceImpl {
	return &customerServiceImpl{customerRepository: customerRepository}
}

func (service *customerServiceImpl) RegisterCustomer(name, email string) (*uuid.UUID, error) {
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

func (service *customerServiceImpl) GetAllCustomers() ([]CustomerDetail, error) {
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

func (service *customerServiceImpl) TopUp(customerId uuid.UUID, amount float64) error {
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
