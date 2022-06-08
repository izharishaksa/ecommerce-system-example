package kafkaclient

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"inventory-service/internal/use_case"
	"log"
)

type Handler interface {
	PlacedOrder(message kafka.Message) error
}

type handler struct {
	inventoryService use_case.InventoryService
}

func NewHandler(inventoryService use_case.InventoryService) Handler {
	return &handler{inventoryService: inventoryService}
}

func (h handler) PlacedOrder(message kafka.Message) error {
	log.Printf("Received message: %s", message.Value)
	var request use_case.PlacedOrderRequest
	err := json.Unmarshal(message.Value, &request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = h.inventoryService.OrderPlaced(request)
	log.Println(err)
	return err
}
