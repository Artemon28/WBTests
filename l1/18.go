package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type inc struct {
	x int64
}

func (i *inc) increment() {
	for {
		time.Sleep(time.Millisecond * 10)
		i.x++
	}
}

func (i *inc) write() {
	fmt.Println(i.x)
}

func main() {
	var i inc
	go i.increment()
	defer i.write()
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT)
	<-shutdownSignal
}
