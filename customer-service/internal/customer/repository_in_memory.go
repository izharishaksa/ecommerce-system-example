package customer

import (
	"fmt"
	"github.com/google/uuid"
	"lib"
)

type inMemoryRepository struct {
	customers map[uuid.UUID]*Customer
}

func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{
		customers: make(map[uuid.UUID]*Customer),
	}
}

func (repo inMemoryRepository) SaveCustomer(customer *Customer) error {
	for _, c := range repo.customers {
		if c.Email == customer.Email {
			return lib.NewErrBadRequest(fmt.Sprintf("customer with email %s already exists", customer.Email))
		}
	}
	repo.customers[customer.Id] = customer
	return nil
}

func (repo inMemoryRepository) GetCustomer() ([]Customer, error) {
	customers := make([]Customer, 0)
	for _, customer := range repo.customers {
		customers = append(customers, *customer)
	}
	return customers, nil
}

func (repo inMemoryRepository) FindCustomerById(id uuid.UUID) (*Customer, error) {
	customer, ok := repo.customers[id]
	if !ok {
		return nil, lib.NewErrNotFound("customer not found")
	}
	return customer, nil
}

func (repo inMemoryRepository) UpdateBalance(customer *Customer) error {
	_, ok := repo.customers[customer.Id]
	if !ok {
		return lib.NewErrNotFound("customer not found")
	}
	repo.customers[customer.Id].Balance = customer.Balance
	return nil
}
