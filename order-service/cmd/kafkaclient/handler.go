package kafkaclient

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"order-service/internal/use_case"
)

type kafkaConsumerService interface {
	OrderCreated(request use_case.OrderCreatedRequest) error
	OrderRejected(request use_case.OrderRejectedRequest) error
}

type handlerImpl struct {
	service kafkaConsumerService
}

func NewHandler(service kafkaConsumerService) *handlerImpl {
	return &handlerImpl{service: service}
}

func (h handlerImpl) OrderRejected(message kafka.Message) error {
	var request use_case.OrderRejectedRequest
	err := json.Unmarshal(message.Value, &request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = h.service.OrderRejected(request)
	return err
}

func (h handlerImpl) OrderCreated(message kafka.Message) error {
	var request use_case.OrderCreatedRequest
	err := json.Unmarshal(message.Value, &request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = h.service.OrderCreated(request)
	return err
}
