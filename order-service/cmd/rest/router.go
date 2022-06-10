package rest

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/segmentio/kafka-go"
	"lib"
	"log"
	"net"
	"net/http"
	"order-service/internal/use_case"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type Handler interface {
	CreateOrder(http.ResponseWriter, *http.Request)
	GetOrders(writer http.ResponseWriter, request *http.Request)
}

func Run(ctx context.Context, cfg lib.Config, requestHandler Handler) error {
	var err error
	err = setKafkaTopic(cfg)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/orders", requestHandler.CreateOrder).Methods("POST")
	router.HandleFunc("/api/v1/orders", requestHandler.GetOrders).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"POST", "GET", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Mode"},
		MaxAge:             60, // 1 minutes
		AllowCredentials:   true,
		OptionsPassthrough: false,
		Debug:              false,
	})
	httpHandler := c.Handler(router)

	err = startServer(ctx, httpHandler, cfg)
	if err != nil {
		return err
	}
	return nil
}

func setKafkaTopic(cfg lib.Config) error {
	conn, err := kafka.Dial("tcp", cfg.Kafka)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{Topic: use_case.OrderStatusPlacedTopic, NumPartitions: 1, ReplicationFactor: 1},
		{Topic: use_case.OrderStatusCreatedTopic, NumPartitions: 1, ReplicationFactor: 1},
		{Topic: use_case.OrderStatusCanceledTopic, NumPartitions: 1, ReplicationFactor: 1},
		{Topic: use_case.OrderStatusPaidTopic, NumPartitions: 1, ReplicationFactor: 1},
		{Topic: use_case.OrderStatusRejectedTopic, NumPartitions: 1, ReplicationFactor: 1},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}

	return nil
}

func startServer(ctx context.Context, httpHandler http.Handler, cfg lib.Config) error {
	errChan := make(chan error, 1)

	go func() {
		errChan <- startHTTP(ctx, httpHandler, cfg)
	}()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func startHTTP(ctx context.Context, httpHandler http.Handler, cfg lib.Config) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.App.HTTPPort),
		Handler: httpHandler,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("%s is running at port %d:", cfg.App.Name, cfg.App.HTTPPort)

	interruption := make(chan os.Signal, 1)

	signal.Notify(interruption, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-interruption

	if err := server.Shutdown(ctx); err != nil {
		return err
	}
	log.Printf("%s is shutting down...", cfg.App.Name)

	return nil
}
