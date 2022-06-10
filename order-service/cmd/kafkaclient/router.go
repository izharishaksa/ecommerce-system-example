package kafkaclient

import (
	"context"
	"github.com/segmentio/kafka-go"
	"lib"
	"log"
)

type Handler interface {
	OrderRejected(message kafka.Message) error
	OrderCreated(message kafka.Message) error
}

func RunConsumer(ctx context.Context, cfg lib.Config, consumerHandler Handler) error {
	errChan := make(chan error, 3)
	go func() {
		errChan <- Consume(ctx, cfg, "ORDER_CREATED", "ORDER_CREATED_GROUP", consumerHandler.OrderCreated)
	}()
	go func() {
		errChan <- Consume(ctx, cfg, "ORDER_REJECTED", "ORDER_REJECTED_GROUP", consumerHandler.OrderRejected)
	}()
	go func() {
		errChan <- Consume(ctx, cfg, "ORDER_PAID", "ORDER_PAID_GROUP", exampleHandler)
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
