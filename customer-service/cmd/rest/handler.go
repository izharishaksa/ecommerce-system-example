package main

import (
	"customer-service/internal/use_case"
	"encoding/json"
	"github.com/google/uuid"
	"lib"
	"net/http"
)

type CustomerService interface {
	RegisterCustomer(name string) (*uuid.UUID, error)
	GetAllCustomers() ([]use_case.CustomerDetail, error)
	TopUp(customerId uuid.UUID, amount float64) error
}

type Handler struct {
	customerService CustomerService
}

func NewHandler(customerService CustomerService) *Handler {
	return &Handler{customerService: customerService}
}

func (h Handler) registerCustomer(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Name string `json:"name"`
	}
	var request requestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), http.StatusBadRequest, nil)
		return
	}
	productId, err := h.customerService.RegisterCustomer(request.Name)
	if err != nil {
		lib.WriteResponse(w, err, http.StatusBadRequest, nil)
		return
	}

	lib.WriteResponse(w, err, http.StatusCreated, productId)
}

func (h Handler) getCustomer(w http.ResponseWriter, _ *http.Request) {
	customers, err := h.customerService.GetAllCustomers()
	if err != nil {
		lib.WriteResponse(w, err, http.StatusInternalServerError, nil)
		return
	}
	lib.WriteResponse(w, err, http.StatusOK, customers)
}

func (h Handler) topUpBalance(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		CustomerId uuid.UUID `json:"customer_id"`
		Amount     float64   `json:"amount"`
	}
	var request requestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), http.StatusBadRequest, nil)
		return
	}
	err = h.customerService.TopUp(request.CustomerId, request.Amount)
	if err != nil {
		lib.WriteResponse(w, err, http.StatusBadRequest, nil)
		return
	}
	lib.WriteResponse(w, err, http.StatusOK, nil)
}
