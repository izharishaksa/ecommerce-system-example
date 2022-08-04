package kafkaclient

import (
	"context"
	"github.com/segmentio/kafka-go"
	"lib"
)

type Handler interface {
	OrderRejected(message kafka.Message) error
	Order(message kafka.Message) error
}

func RunConsumer(ctx context.Context, cfg lib.Config, consumerHandler Handler) error {
	errChan := make(chan error, 1)
	go func() {
		errChan <- Consume(ctx, cfg, "ORDER", "ORDER_GROUP", consumerHandler.Order)
	}()
	return <-errChan
}
