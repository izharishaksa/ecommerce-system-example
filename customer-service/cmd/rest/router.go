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
	RegisterCustomer(http.ResponseWriter, *http.Request)
	GetCustomer(http.ResponseWriter, *http.Request)
	TopUpBalance(http.ResponseWriter, *http.Request)
}

func Run(ctx context.Context, cfg lib.Config, requestHandler Handler) error {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/customers", requestHandler.RegisterCustomer).Methods("POST")
	router.HandleFunc("/api/v1/customers", requestHandler.GetCustomer).Methods("GET")
	router.HandleFunc("/api/v1/top_up", requestHandler.TopUpBalance).Methods("POST")

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

	err := startServer(ctx, httpHandler, cfg)
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
			log.Fatal("failed to start server: ", err)
		}
	}()
	log.Printf("%s is starting at port %d:", cfg.App.Name, cfg.App.HTTPPort)

	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-interruption

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("failed to shutdown server: %s", err.Error())
		return err
	}
	log.Printf("%s is shutting down gracefully...", cfg.App.Name)

	return nil
}
