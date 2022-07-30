package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var x float64
	decades := make(map[int][]float64)
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT)
	go func() {
		for {
			_, err := fmt.Scan(&x)
			if err != nil {
				log.Fatal(err.Error())
			}
			y := int(x / 10)
			if y > 0 || y == 0 && x > 0 {
				y++
			} else {
				y--
			}
			decades[y] = append(decades[y], x)
		}
	}()
	<-shutdownSignal
	for key, val := range decades {
		fmt.Printf("%d: ", key*10)
		fmt.Println(val)
	}
}
