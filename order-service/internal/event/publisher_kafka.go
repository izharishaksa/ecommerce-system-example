package event

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type kafkaPublisher struct {
	producer *kafka.Writer
}

func NewKafkaPublisher(producer *kafka.Writer) *kafkaPublisher {
	return &kafkaPublisher{producer: producer}
}

func (k kafkaPublisher) Publish(ctx context.Context, topic string, id string, value []byte) error {
	err := k.producer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(id),
		Value: value,
		Topic: topic,
	})
	if err != nil {
		return err
	}
	return nil
}
