package use_case

import (
	"customer-service/internal/customer"
	"github.com/google/uuid"
	"sync"
)

type customerRepository interface {
	SaveCustomer(customer *customer.Customer) error
	GetCustomer() ([]customer.Customer, error)
	FindCustomerById(id uuid.UUID) (*customer.Customer, error)
	UpdateBalance(customer *customer.Customer) error
}

type CustomerService struct {
	customerRepository customerRepository
	mutexes            map[uuid.UUID]*sync.Mutex
	mutexesMu          sync.Mutex
}

func NewCustomerService(customerRepository customerRepository) *CustomerService {
	return &CustomerService{
		customerRepository: customerRepository,
		mutexes:            make(map[uuid.UUID]*sync.Mutex),
	}
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
		customerDetails = append(customerDetails, fromCustomerToCustomerDetail(cust))
	}
	return customerDetails, nil
}

func (service *CustomerService) TopUp(customerId uuid.UUID, amount float64) error {
	mu := service.getMutex(customerId)
	mu.Lock()
	defer mu.Unlock()

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

func (service *CustomerService) getMutex(customerId uuid.UUID) *sync.Mutex {
	service.mutexesMu.Lock()
	defer service.mutexesMu.Unlock()

	mu, exists := service.mutexes[customerId]
	if !exists {
		mu = &sync.Mutex{}
		service.mutexes[customerId] = mu
	}

	return mu
}
