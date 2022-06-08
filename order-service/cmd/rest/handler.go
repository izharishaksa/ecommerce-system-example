package rest

import (
	"encoding/json"
	"lib"
	"net/http"
	"order-service/internal/use_case"
)

type Handler interface {
	CreateOrder(http.ResponseWriter, *http.Request)
	GetOrders(writer http.ResponseWriter, request *http.Request)
}

type handler struct {
	orderService use_case.OrderService
}

func NewHandler(orderService use_case.OrderService) Handler {
	return &handler{orderService: orderService}
}

func (h handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var request use_case.CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	productId, err := h.orderService.CreateOrder(request)

	lib.WriteResponse(w, err, productId)
}

func (h handler) GetOrders(writer http.ResponseWriter, _ *http.Request) {
	orders, err := h.orderService.GetAllOrders()
	lib.WriteResponse(writer, err, orders)
}
