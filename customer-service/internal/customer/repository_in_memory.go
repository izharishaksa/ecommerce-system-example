package customer

import (
	"github.com/google/uuid"
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
	repo.customers[customer.Id] = customer
	return nil
}
