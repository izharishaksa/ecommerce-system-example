package main

import (
	"lib"
	"log"
	"order-service/cmd/rest"
	"sync"
)

func main() {
	cfg := lib.LoadConfigByFile("./cmd", "config", "yaml")

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		err := rest.Run(cfg)
		if err != nil {
			log.Println(err)
		}
		wg.Done()
	}()

	wg.Wait()
}
