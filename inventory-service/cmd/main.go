package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"inventory-service/cmd/kafkaclient"
	"inventory-service/cmd/rest"
	"inventory-service/internal/inventory"
	"inventory-service/internal/use_case"
	"lib"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func exampleHandler(message kafka.Message) error {
	log.Printf("message: %s %s", message.Topic, message.Value)

	return nil
}

func main() {
	cfg := lib.LoadConfigByFile("./cmd", "config", "yaml")
	var ctx, cancel = context.WithCancel(context.Background())

	//setup service
	kafkaWriter := &kafka.Writer{
		Addr:                   kafka.TCP(cfg.Kafka),
		AllowAutoTopicCreation: true,
	}
	inventoryRepository := inventory.NewInMemoryRepository()
	inventoryService := use_case.NewInventoryService(inventoryRepository, kafkaWriter)

	//setup rest handler
	restChan := make(chan error, 1)
	go func() {
		restHandler := rest.NewHandler(inventoryService)
		restChan <- rest.RunServer(cfg, restHandler)
	}()

	//setup kafka consumer handler
	topics := []string{"ORDER_PLACED", "ORDER_CREATED", "ORDER_CANCELED"}
	consumerErrChan := make(chan error, len(topics))
	kafkaConsumerHandler := kafkaclient.NewHandler(inventoryService)
	go func() {
		consumerErrChan <- kafkaclient.Consume(ctx, cfg, "ORDER_PLACED", "ORDER_PLACED_GROUP", kafkaConsumerHandler.PlacedOrder)
	}()
	go func() {
		consumerErrChan <- kafkaclient.Consume(ctx, cfg, "ORDER_CREATED", "ORDER_CREATED_GROUP", exampleHandler)
	}()
	go func() {
		consumerErrChan <- kafkaclient.Consume(ctx, cfg, "ORDER_CANCELED", "ORDER_CANCELED_GROUP", exampleHandler)
	}()

	interruption := make(chan os.Signal)
	go func() {
		signal.Notify(interruption, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
	}()

	<-interruption
	cancel()

	select {
	case <-interruption:
		log.Println("Interrupted")
	case err := <-consumerErrChan:
		log.Println("consumer ran with an error", err)
	case err := <-restChan:
		log.Println("rest ran with an error", err)
	}
}
