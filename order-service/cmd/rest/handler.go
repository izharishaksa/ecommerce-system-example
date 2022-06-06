package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"lib"
	"net/http"
	"order-service/internal/use_case"
)

type OrderService interface {
	CreateOrder(request use_case.CreateOrderRequest) (*uuid.UUID, error)
}

type Handler struct {
	orderService OrderService
}

func NewHandler(orderService OrderService) *Handler {
	return &Handler{orderService: orderService}
}

func (h Handler) createOrder(w http.ResponseWriter, r *http.Request) {
	var request use_case.CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	productId, err := h.orderService.CreateOrder(request)

	lib.WriteResponse(w, err, productId)
}
