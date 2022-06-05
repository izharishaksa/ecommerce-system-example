package customer

import (
	"github.com/google/uuid"
	"lib"
	"strings"
)

type Customer struct {
	Id      uuid.UUID
	Name    string
	Balance float64
}

func NewCustomer(name string) (*Customer, error) {
	instance := &Customer{
		Id: uuid.New(),
	}
	err := instance.SetName(name)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (c *Customer) SetName(name string) error {
	if strings.TrimSpace(name) == "" {
		return lib.NewErrBadRequest("name cannot be empty")
	}
	c.Name = name
	return nil
}

func (c *Customer) TopUp(amount float64) error {
	if amount <= 0 {
		return lib.NewErrBadRequest("amount must be greater than 0")
	}
	c.Balance += amount
	return nil
}
