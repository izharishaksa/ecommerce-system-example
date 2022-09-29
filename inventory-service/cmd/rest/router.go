package rest

import (
	"context"
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

type Handler interface {
	CreateProduct(http.ResponseWriter, *http.Request)
	GetProduct(http.ResponseWriter, *http.Request)
	AddProductStock(http.ResponseWriter, *http.Request)
}

func RunServer(cfg lib.Config, requestHandler Handler) error {
	ctx := context.TODO()
	var err error

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/products", requestHandler.GetProduct).Methods("GET")
	router.HandleFunc("/api/v1/products", requestHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/api/v1/stocks", requestHandler.AddProductStock).Methods("POST")

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
			log.Fatalf("%s failed to start: %v", cfg.App.Name, err)
		}
	}()
	log.Printf("%s is starting at port %d:", cfg.App.Name, cfg.App.HTTPPort)

	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-interruption

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("%s failed to shutdown: %v", cfg.App.Name, err)
		return err
	}
	log.Printf("%s is shutting down...", cfg.App.Name)

	return nil
}
