package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func fillChan(ch chan int, array []int) {
	for _, v := range array {
		ch <- v
	}
}

func fillSquareChan(values chan int, results chan int) {
	defer close(values)
	for v := range values {
		results <- v * v
	}
}

func print(results chan int) {
	for v := range results {
		fmt.Println(v)
	}
}

func main() {
	array := []int{1, 2, 4, 6, 124, 64, 12, 312, 6, 4, 2, 412, 3, 2, 64, 12, 4, 2, 5, 6, 4, 12, 4, 1, 2, 4, 6, 124, 64, 12, 312, 6, 4, 2, 412, 3, 2, 64, 12, 4, 2, 5, 6, 4, 12, 4, 1, 2, 4, 6, 124, 64, 12, 312, 6, 4, 2, 412, 3, 2, 64, 12, 4, 2, 5, 6, 4, 12, 4, 1, 2, 4, 6, 124, 64, 12, 312, 6, 4, 2, 412, 3, 2, 64, 12, 4, 2, 5, 6, 4, 12, 4, 1, 2, 4, 6, 124, 64, 12, 312, 6, 4, 2, 412, 3, 2, 64, 12, 4, 2, 5, 6, 4, 12, 4, 1, 2, 4, 6, 124, 64, 12, 312, 6, 4, 2, 412, 3, 2, 64, 12, 4, 2, 5, 6, 4, 12, 4, 1, 2, 4, 6, 124, 64, 12, 312, 6, 4, 2, 412, 3, 2, 64, 12, 4, 2, 5, 6, 4, 12, 4, 1, 2, 4, 6, 124, 64, 12, 312, 6, 4, 2, 412, 3, 2, 64, 12, 4, 2, 5, 6, 4, 12, 4, 1, 2, 4, 6, 124, 64, 12, 312, 6, 4, 2, 412, 3, 2, 64, 12, 4, 2, 5, 6, 4, 12, 4, 1, 2, 4, 6, 124, 64, 12, 312, 6, 4, 2, 412, 3, 2, 64, 12, 4, 2, 5, 6, 4, 12, 4, 1, 2, 4, 6, 124, 64, 12, 312, 6, 4, 2, 412, 3, 2, 64, 12, 4, 2, 5, 6, 4, 12, 4, 1, 2, 4, 6, 124, 64, 12, 312, 6, 4, 2, 412, 3, 2, 64, 12, 4, 2, 5, 6, 4, 12, 4, 1, 2, 4, 6, 124, 64, 12, 312, 6, 4, 2, 412, 3, 2, 64, 12, 4, 2, 5, 6, 4, 12, 4}
	values := make(chan int)
	squares := make(chan int)
	go fillChan(values, array)
	go fillSquareChan(values, squares)
	go print(squares)

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT)
	<-shutdownSignal
	close(squares)
}
