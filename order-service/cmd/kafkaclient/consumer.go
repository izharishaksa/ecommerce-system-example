package kafkaclient

import (
	"context"
	"github.com/segmentio/kafka-go"
	"lib"
	"log"
)

func Consume(ctx context.Context, cfg lib.Config, topic string, consumerGroup string, handler func(message kafka.Message) error) error {
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.Kafka},
		GroupID: consumerGroup,
		Topic:   topic,
	})
	defer kafkaReader.Close()

	log.Printf("Start consuming topic: %s", topic)
	errChan := make(chan error, 1)
	for {
		msg, err := kafkaReader.ReadMessage(ctx)
		if err != nil {
			errChan <- err
			break
		}

		if err := handler(msg); err != nil {
			log.Printf("error consuming message, err: %#v\n", err)
			continue
		}
	}

	err := <-errChan
	log.Printf("consumer stopped with an error %#v\n", err.Error())
	if err != nil {
		panic(err)
	}
	return err
}
