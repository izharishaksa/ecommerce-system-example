package kafkaclient

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"order-service/internal/use_case"
)

type Handler interface {
	OrderRejected(message kafka.Message) error
	OrderCreated(message kafka.Message) error
}

type handler struct {
	orderService use_case.OrderService
}

func NewHandler(orderService use_case.OrderService) Handler {
	return &handler{orderService: orderService}
}

func (h handler) OrderRejected(message kafka.Message) error {
	var request use_case.OrderRejectedRequest
	err := json.Unmarshal(message.Value, &request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = h.orderService.OrderRejected(request)
	return err
}

func (h handler) OrderCreated(message kafka.Message) error {
	var request use_case.OrderCreatedRequest
	err := json.Unmarshal(message.Value, &request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = h.orderService.OrderCreated(request)
	return err
}
