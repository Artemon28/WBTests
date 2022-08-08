package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func reader(chanToRead <-chan int, ctx context.Context) {
	for {
		select {
		case i := <-chanToRead:
			fmt.Println(i)
		case <-ctx.Done():
			return
		}

	}
}

func writer(chanToWrite chan<- int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			chanToWrite <- rand.Int()
		}

	}
}

func main() {
	var seconds string
	fmt.Println("Enter amount of seconds")
	fmt.Scan(&seconds)
	seconds += "s"
	d, _ := time.ParseDuration(seconds)
	ch := make(chan int)

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, d)

	go writer(ch, ctx)
	go reader(ch, ctx)

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT)
	<-shutdownSignal
}
