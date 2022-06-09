package kafkaclient

import "github.com/segmentio/kafka-go"

type Handler interface {
	PlacedOrder(kafka.Message) error
}
