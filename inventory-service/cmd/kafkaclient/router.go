package kafkaclient

import (
	"context"
	"github.com/segmentio/kafka-go"
	"lib"
	"log"
)

type handler interface {
	PlacedOrder(kafka.Message) error
}

func RunConsumer(ctx context.Context, cfg lib.Config, consumerHandler handler) error {
	errChan := make(chan error, 3)
	go func() {
		errChan <- Consume(ctx, cfg, "ORDER_PLACED", "ORDER_PLACED_GROUP", consumerHandler.PlacedOrder)
	}()
	go func() {
		errChan <- Consume(ctx, cfg, "ORDER_CREATED", "ORDER_CREATED_GROUP", exampleHandler)
	}()
	go func() {
		errChan <- Consume(ctx, cfg, "ORDER_CANCELED", "ORDER_CANCELED_GROUP", exampleHandler)
	}()
	return <-errChan
}

func exampleHandler(message kafka.Message) error {
	log.Printf("message: %s %s", message.Topic, message.Value)

	return nil
}
