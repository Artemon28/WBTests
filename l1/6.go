package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func close1(quit <-chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			continue
		}
	}
}

func close2(ch <-chan string, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for i := range ch {
		fmt.Println(i)
	}
}

func close3(ctx context.Context, str string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(str)
		}
	}
}

func main() {

	//close 1
	quit := make(chan bool)
	go close1(quit)

	quit <- true

	//close 2
	wg := sync.WaitGroup{}
	chan2 := make(chan string)
	go close2(chan2, &wg)
	chan2 <- "hello"
	chan2 <- "bye"
	close(chan2)
	wg.Wait()

	//close 3
	ctx, cancel := context.WithCancel(context.Background())
	go close3(ctx, "test1")

	//close 4
	ctx2 := context.Background()
	ctx2, _ = context.WithTimeout(ctx2, time.Second*5)
	go close3(ctx2, "test2")

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT)
	<-shutdownSignal
	cancel()
}
