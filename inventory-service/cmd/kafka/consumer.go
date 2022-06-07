package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"golang.org/x/sync/errgroup"
	"lib"
	"log"
	"sync"
)

func Consume(ctx context.Context, cfg lib.Config, topics []string) {
	g, _ := errgroup.WithContext(ctx)
	wg := new(sync.WaitGroup)
	for _, t := range topics {
		topic := t
		wg.Add(1)
		g.Go(func() error {
			kafkaReader := kafka.NewReader(kafka.ReaderConfig{
				Brokers: []string{cfg.Kafka},
				GroupID: "1",
				Topic:   topic,
			})
			defer func(kafkaReader *kafka.Reader) {
				err := kafkaReader.Close()
				if err != nil {
					log.Println("Error closing kafka reader: ", err)
				}
			}(kafkaReader)
			fmt.Println("Start read message from topic: ", topic)
			for {
				msg, err := kafkaReader.ReadMessage(ctx)
				if err != nil {
					return err
				}
				log.Printf("%s %s %s", msg.Value, msg.Key, topic)
			}
		})
	}
	err := g.Wait()
	if err != nil {
		fmt.Println("Error kafka group: ", err)
	}
	fmt.Println("Stop reading all kafka messages")
}
