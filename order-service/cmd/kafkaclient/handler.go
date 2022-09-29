package kafkaclient

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"order-service/internal/order"
)

type kafkaConsumerService interface {
	OrderCreated(request order.OrderPlacedMessage) error
	OrderRejected(request order.OrderRejectedMessage) error
}

type handlerImpl struct {
	service kafkaConsumerService
}

func NewHandler(service kafkaConsumerService) *handlerImpl {
	return &handlerImpl{service: service}
}

func (h handlerImpl) OrderRejected(message kafka.Message) error {
	var request order.OrderRejectedMessage
	err := json.Unmarshal(message.Value, &request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = h.service.OrderRejected(request)
	return err
}

func (h handlerImpl) Order(message kafka.Message) error {
	var request order.OrderPlacedMessage
	err := json.Unmarshal(message.Value, &request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = h.service.OrderCreated(request)
	return err
}
