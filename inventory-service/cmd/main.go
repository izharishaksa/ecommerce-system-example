package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"inventory-service/cmd/kafkaclient"
	"inventory-service/cmd/rest"
	"inventory-service/internal/event"
	"inventory-service/internal/inventory"
	"inventory-service/internal/use_case"
	"lib"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := lib.LoadConfigByFile("./cmd", "config", "yaml")
	var ctx, cancel = context.WithCancel(context.Background())

	//setup service
	kafkaWriter := &kafka.Writer{
		Addr:                   kafka.TCP(cfg.Kafka),
		AllowAutoTopicCreation: true,
	}
	inventoryRepository := inventory.NewInMemoryRepository()
	eventPublisher := event.NewKafkaPublisher(kafkaWriter)
	inventoryService := use_case.NewInventoryService(inventoryRepository, eventPublisher)

	//setup rest server
	restChan := make(chan error, 1)
	go func() {
		restHandler := rest.NewHandler(inventoryService)
		restChan <- rest.RunServer(cfg, restHandler)
	}()

	//setup kafka consumer
	consumerErrChan := make(chan error, 1)
	kafkaConsumerHandler := kafkaclient.NewHandler(inventoryService)
	go func() {
		consumerErrChan <- kafkaclient.RunConsumer(ctx, cfg, kafkaConsumerHandler)
	}()

	interruption := make(chan os.Signal)
	go func() {
		signal.Notify(interruption, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
	}()

	<-interruption
	cancel()

	select {
	case <-interruption:
		log.Println("interrupted")
	case err := <-consumerErrChan:
		log.Printf("consumer error: %s", err.Error())
	case err := <-restChan:
		log.Println("rest ran with an error", err)
	}
}
