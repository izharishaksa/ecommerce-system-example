package main

import (
	"inventory-service/cmd/rest"
	"lib"
	"log"
	"sync"
)

func main() {
	cfg := lib.LoadConfigByFile("./cmd", "config", "yaml")
	cfg1 := lib.LoadConfigByFile("./cmd", "config", "yaml")

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		err := rest.Run(cfg)
		if err != nil {
			log.Println(err)
		}
		wg.Done()
	}()

	go func() {
		cfg1.App.HTTPPort = 6001
		err := rest.Run(cfg1)
		if err != nil {
			log.Println(err)
		}
		wg.Done()
	}()

	wg.Wait()
}
