package customer

import (
	"fmt"
	"github.com/google/uuid"
	"lib"
)

type InMemoryRepository struct {
	customers map[uuid.UUID]*Customer
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		customers: make(map[uuid.UUID]*Customer),
	}
}

func (repo InMemoryRepository) SaveCustomer(customer *Customer) error {
	for _, c := range repo.customers {
		if c.Email == customer.Email {
			return lib.NewErrBadRequest(fmt.Sprintf("customer with email %s already exists", customer.Email))
		}
	}
	repo.customers[customer.Id] = customer
	return nil
}

func (repo InMemoryRepository) GetCustomer() ([]Customer, error) {
	customers := make([]Customer, 0)
	for _, customer := range repo.customers {
		customers = append(customers, *customer)
	}
	return customers, nil
}

func (repo InMemoryRepository) FindCustomerById(id uuid.UUID) (*Customer, error) {
	customer, ok := repo.customers[id]
	if !ok {
		return nil, lib.NewErrNotFound("customer not found")
	}
	return customer, nil
}
