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
