package kafkaclient

import (
	"context"
	"github.com/segmentio/kafka-go"
	"lib"
	"log"
)

func Consume(ctx context.Context, cfg lib.Config, topic string, consumerGroup string, handler func(message kafka.Message) error) error {
	errChan := make(chan error, 1)

	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.Kafka},
		GroupID: consumerGroup,
		Topic:   topic,
	})
	defer func(kafkaReader *kafka.Reader) {
		err := kafkaReader.Close()
		if err != nil {
			errChan <- err
		}
	}(kafkaReader)

	for {
		msg, err := kafkaReader.ReadMessage(ctx)
		if err != nil {
			errChan <- err
			break
		}

		log.Printf("Received message: %s", msg.Value)
		if err := handler(msg); err != nil {
			log.Printf("error consuming message, err: %#v\n", err)
		}
	}

	err := <-errChan
	log.Printf("consumer %s stopped with an error %#v\n", topic, err.Error())
	if err != nil {
		panic(err)
	}
	return err
}
