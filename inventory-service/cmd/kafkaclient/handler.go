package kafkaclient

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"inventory-service/internal/use_case"
	"log"
)

type eventConsumerService interface {
	OrderPlaced(request use_case.PlacedOrderRequest) error
}

type handlerImpl struct {
	service eventConsumerService
}

func NewHandler(service eventConsumerService) *handlerImpl {
	return &handlerImpl{service: service}
}

func (h handlerImpl) PlacedOrder(message kafka.Message) error {
	log.Printf("Received message: %s", message.Value)
	var request use_case.PlacedOrderRequest
	err := json.Unmarshal(message.Value, &request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = h.service.OrderPlaced(request)
	log.Println(err)

	return err
}
