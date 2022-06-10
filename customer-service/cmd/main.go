package main

import (
	"context"
	"customer-service/cmd/rest"
	"customer-service/internal/customer"
	"customer-service/internal/use_case"
	"lib"
	"log"
	"sync"
)

func main() {
	cfg := lib.LoadConfigByFile("./cmd", "config", "yaml")
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		ctx := context.Background()
		customerRepository := customer.NewInMemoryRepository()
		customerService := use_case.NewCustomerService(customerRepository)
		requestHandler := rest.NewHandler(customerService)

		err := rest.Run(ctx, cfg, requestHandler)
		if err != nil {
			log.Println(err)
		}
		wg.Done()
	}()
	wg.Wait()
}
