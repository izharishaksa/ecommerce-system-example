package rest

import (
	"customer-service/internal/use_case"
	"encoding/json"
	"github.com/google/uuid"
	"lib"
	"net/http"
)

type Handler interface {
	RegisterCustomer(http.ResponseWriter, *http.Request)
	GetCustomer(http.ResponseWriter, *http.Request)
	TopUpBalance(http.ResponseWriter, *http.Request)
}

type handler struct {
	customerService use_case.CustomerService
}

func NewHandler(customerService use_case.CustomerService) Handler {
	return &handler{customerService: customerService}
}

func (h handler) RegisterCustomer(w http.ResponseWriter, r *http.Request) {
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

func (h handler) GetCustomer(w http.ResponseWriter, _ *http.Request) {
	customers, err := h.customerService.GetAllCustomers()
	if err != nil {
		lib.WriteResponse(w, err, nil)
		return
	}
	lib.WriteResponse(w, err, customers)
}

func (h handler) TopUpBalance(w http.ResponseWriter, r *http.Request) {
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
