package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"lib"
	"net/http"
)

type CustomerService interface {
	RegisterCustomer(name string) (*uuid.UUID, error)
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

	lib.WriteResponse(w, err, http.StatusCreated, productId)
}
