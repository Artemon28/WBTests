package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func worker(values <-chan string, quit <-chan bool) {
	for {
		select {
		case v := <-values:
			fmt.Println(v)
		case <-quit:
			return
		}
	}
}

func writeToStdin2(strs chan<- string) {
	for {
		var str string
		fmt.Scan(&str)
		strs <- str
	}
}

// но зачем нам ждать, когда горутины дочитают из канала, если нам необходимо выйти прямо сейчас?
// Этот вариант будет слушать канал и закрываться, когда приходит. Если же горутины выполняются долго,
//то верным решением будет быстрая остановка каждой из них через канал оповещения.
func main() {
	values := make(chan string)
	quit := make(chan bool)
	var n int
	fmt.Println("Insert amount of workers")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err.Error())
	}
	for i := 1; i <= n; i++ {
		go worker(values, quit)
	}

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT)

	fmt.Println("Enter your strings here:")
	go writeToStdin2(values)

	<-shutdownSignal
	quit <- true
	close(values)
}
