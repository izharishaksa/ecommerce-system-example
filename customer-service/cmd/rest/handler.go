package rest

import (
	"customer-service/internal/use_case"
	"encoding/json"
	"github.com/google/uuid"
	"lib"
	"net/http"
)

type CustomerService interface {
	RegisterCustomer(name, email string) (*uuid.UUID, error)
	GetAllCustomers() ([]use_case.CustomerDetail, error)
	TopUp(customerId uuid.UUID, amount float64) error
}

type handlerImpl struct {
	customerService CustomerService
}

func NewHandler(customerService CustomerService) *handlerImpl {
	return &handlerImpl{customerService: customerService}
}

func (h handlerImpl) RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	var request requestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	productId, err := h.customerService.RegisterCustomer(request.Name, request.Email)
	if err != nil {
		lib.WriteResponse(w, err, nil)
		return
	}

	lib.WriteResponse(w, err, productId)
}

func (h handlerImpl) GetCustomer(w http.ResponseWriter, _ *http.Request) {
	customers, err := h.customerService.GetAllCustomers()
	if err != nil {
		lib.WriteResponse(w, err, nil)
		return
	}
	lib.WriteResponse(w, err, customers)
}

func (h handlerImpl) TopUpBalance(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		CustomerId uuid.UUID `json:"customer_id"`
		Amount     float64   `json:"amount"`
	}
	var request requestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	err = h.customerService.TopUp(request.CustomerId, request.Amount)
	if err != nil {
		lib.WriteResponse(w, err, nil)
		return
	}
	lib.WriteResponse(w, err, nil)
}
