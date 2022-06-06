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

func (service *CustomerService) RegisterCustomer(name, email string) (*uuid.UUID, error) {
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

func (service *CustomerService) GetAllCustomers() ([]CustomerDetail, error) {
	customers, err := service.customerRepository.GetCustomer()
	if err != nil {
		return nil, err
	}
	customerDetails := make([]CustomerDetail, 0)
	for _, cust := range customers {
		customerDetails = append(customerDetails, CustomerDetail{
			ID:      cust.Id,
			Name:    cust.Name,
			Balance: cust.Balance,
		})
	}
	return customerDetails, nil
}

func (service *CustomerService) TopUp(customerId uuid.UUID, amount float64) error {
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
