package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"inventory-service/cmd/kafka"
	"inventory-service/cmd/rest"
	"lib"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func iterate(ctx context.Context) {
	g, _ := errgroup.WithContext(ctx)
	for ii := 5; ii < 6; ii++ {
		i := ii
		g.Go(func() error {
			fmt.Println("Looping...", i)
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Break the loop ", i)
					return nil
				case <-time.After(time.Duration(i) * time.Second):
					fmt.Println("Hello in a loop ", i)
				}
			}
		})
	}
	err := g.Wait()
	if err != nil {
		fmt.Println("Error group: ", err)
	}
	fmt.Println("inside iteration done")
}

func main() {
	cfg := lib.LoadConfigByFile("./cmd", "config", "yaml")
	var err error
	var ctx, cancel = context.WithCancel(context.Background())
	go func() {
		interruption := make(chan os.Signal, 1)
		signal.Notify(interruption, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
		<-interruption
		cancel()
	}()

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		err = rest.Run(cfg)
		if err != nil {
			log.Println(err)
		}
		return err
	})

	g.Go(func() error {
		topics := []string{"ORDER_PLACED", "ORDER_CREATED", "ORDER_CANCELED", "ORDER_REJECTED", "ORDER_PAID"}
		kafka.Consume(ctx, cfg, topics)
		return nil
	})

	err = g.Wait()
	if err != nil {
		fmt.Println("Error group: ", err)
	}
	fmt.Println("Main done")
}
