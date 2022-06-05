package main

import (
	"context"
	"customer-service/internal/customer"
	"customer-service/internal/use_case"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"lib"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	cfg := lib.LoadConfigByFile("./cmd/rest/", "config", "yaml")
	mysqlConnection, err := lib.NewPostgresqlConnection(cfg.Database)
	if err != nil {
		log.Println(err)
	} else {
		err = mysqlConnection.Ping()
		if err != nil {
			log.Println(err)
		} else {
			log.Println("mysql connection success")
		}
	}

	customerRepository := customer.NewInMemoryRepository()
	customerService := use_case.NewCustomerService(customerRepository)
	handler := NewHandler(customerService)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/customers", handler.registerCustomer).Methods("POST")

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
		log.Fatal(err)
		return
	}
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
	log.Printf("%s is starting at port %d:", cfg.App.Name, cfg.App.HTTPPort)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.App.HTTPPort),
		Handler: httpHandler,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return gracefulShutdown(ctx, server, cfg)
}

func gracefulShutdown(ctx context.Context, server *http.Server, cfg lib.Config) error {
	interruption := make(chan os.Signal, 1)
	defer log.Printf("%s is shutting down...", cfg.App.Name)

	signal.Notify(interruption, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-interruption

	if err := server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
