package kafkaclient

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"order-service/internal/event"
)

type kafkaConsumerService interface {
	OrderCreated(request event.OrderPlacedMessage) error
	OrderRejected(request event.OrderRejectedMessage) error
}

type handlerImpl struct {
	service kafkaConsumerService
}

func NewHandler(service kafkaConsumerService) *handlerImpl {
	return &handlerImpl{service: service}
}

func (h handlerImpl) OrderRejected(message kafka.Message) error {
	var request event.OrderRejectedMessage
	err := json.Unmarshal(message.Value, &request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = h.service.OrderRejected(request)
	return err
}

func (h handlerImpl) Order(message kafka.Message) error {
	var request event.OrderPlacedMessage
	err := json.Unmarshal(message.Value, &request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = h.service.OrderCreated(request)
	return err
}
