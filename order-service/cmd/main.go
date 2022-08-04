package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"lib"
	"log"
	"order-service/cmd/kafkaclient"
	"order-service/cmd/rest"
	"order-service/internal/event"
	"order-service/internal/order"
	"order-service/internal/use_case"
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
	orderRepository := order.NewInMemoryRepository()
	eventPublisher := event.NewKafkaPublisher(kafkaWriter)
	orderService := use_case.NewOrderService(orderRepository, eventPublisher)

	//setup rest handler
	restChan := make(chan error, 1)
	go func() {
		restHandler := rest.NewHandler(orderService)
		restChan <- rest.Run(ctx, cfg, restHandler)
	}()

	//setup kafka consumer handler
	consumerErrChan := make(chan error, 4)
	kafkaConsumerHandler := kafkaclient.NewHandler(orderService)
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
		log.Println("Interrupted")
	case err := <-consumerErrChan:
		log.Println("consumer ran with an error", err)
	case err := <-restChan:
		log.Println("rest ran with an error", err)
	}
}
