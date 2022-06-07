package main

import (
	"context"
	"fmt"
	"lib"
	"log"
	"order-service/cmd/rest"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func check(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return
		case <-time.After(time.Second * 10):
			fmt.Println("Hello")
		}
	}
}

func main() {
	cfg := lib.LoadConfigByFile("./cmd", "config", "yaml")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		interruption := make(chan os.Signal, 1)
		signal.Notify(interruption, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
		<-interruption
		cancel()
	}()

	//conn, err := kafkalib.Dial("tcp", cfg.Kafka)
	//if err != nil {
	//	panic(err.Error())
	//} else {
	//	fmt.Println("Connected to Kafka")
	//}
	//_, err = conn.Brokers()
	//if err != nil {
	//	panic(err.Error())
	//} else {
	//	fmt.Println("Connected to Kafka 1")
	//}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		err := rest.Run(ctx, cfg)
		if err != nil {
			log.Println(err)
		}
		wg.Done()
	}()
	wg.Wait()
}
