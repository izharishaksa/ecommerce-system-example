package rest

import (
	"encoding/json"
	"github.com/google/uuid"
	"lib"
	"net/http"
	"order-service/internal/use_case"
)

type OrderService interface {
	PlaceOrder(request use_case.PlaceOrderRequest) (*uuid.UUID, error)
	GetAllOrders() ([]use_case.OrderResponse, error)
}

type handlerImpl struct {
	orderService OrderService
}

func NewHandler(orderService OrderService) *handlerImpl {
	return &handlerImpl{orderService: orderService}
}

func (h handlerImpl) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var request use_case.PlaceOrderRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	productId, err := h.orderService.PlaceOrder(request)

	lib.WriteResponse(w, err, productId)
}

func (h handlerImpl) GetOrders(writer http.ResponseWriter, _ *http.Request) {
	orders, err := h.orderService.GetAllOrders()
	lib.WriteResponse(writer, err, orders)
}
