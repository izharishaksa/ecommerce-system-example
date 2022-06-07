package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var ctx, cancel = context.WithCancel(context.Background())

func main() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		router := mux.NewRouter()
		router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("Hello from products"))
			fmt.Println("Hello from products")
		}).Methods("GET")
		server := &http.Server{
			Addr:    fmt.Sprintf(":%d", 4080),
			Handler: router,
		}
		go func() {
			fmt.Println("Starting server on port 4080")
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("apps failed to start: %v", err)
			}
		}()

		<-exit

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("failed to shutdown: %v\n", err)
		} else {
			log.Println("Server gracefully stopped")
		}
		return nil
	})

	g.Go(func() error {
		for {
			select {
			case <-time.After(1 * time.Second):
				fmt.Println("Hello in a loop")
			case <-exit:
				fmt.Println("Interrupted the hello loop")
				return nil
			}
		}
	})

	//g.Go(func() error {
	//	for {
	//		select {
	//		case <-gCtx.Done():
	//			fmt.Println("Break the ciao loop")
	//			return nil
	//		case <-time.After(1 * time.Second):
	//			fmt.Println("Ciao in a loop")
	//		}
	//	}
	//})

	err := g.Wait()
	if err != nil {
		fmt.Println("Error group: ", err)
	}

	fmt.Println("Main done")
}

func usingWithGroup() {
	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		select {
		case <-exit:
			fmt.Println("Break the loop from signal")
			cancel()
		}
	}()

	var wg sync.WaitGroup

	counter := 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Break hello loop")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Hello in a loop ", counter)
				counter++
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Break ciao loop")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Ciao in a loop ", counter)
				counter++
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		router := mux.NewRouter()
		router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("Hello from products"))
		}).Methods("GET")
		server := &http.Server{
			Addr:    fmt.Sprintf(":%d", 4080),
			Handler: router,
		}
		go func() {
			fmt.Println("Starting server on port 4080")
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("apps failed to start: %v", err)
			}
		}()

		interruption := make(chan os.Signal, 1)
		defer log.Printf("%s is shutting down gracefully...", "apps")

		signal.Notify(interruption, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
		<-interruption

		select {
		case <-ctx.Done():
			if err := server.Shutdown(ctx); err != nil {
				log.Fatalf("failed to shutdown: %v", err)
			} else {
				log.Println("Server gracefully stopped")
			}
			ctx.Done()
			return
		}
	}()

	wg.Wait()
	fmt.Println("Main done")
}
