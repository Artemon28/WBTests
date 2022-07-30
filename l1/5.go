package main

import (
	"fmt"
	"math/rand"
	"time"
)

func reader(chanToRead <-chan int, quit <-chan bool) {
	for {
		select {
		case i := <-chanToRead:
			fmt.Println(i)
		case <-quit:
			return
		}

	}
}

func writer(chanToWrite chan<- int, quit <-chan bool) {
	for {
		select {
		case <-quit:
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
	quit := make(chan bool)
	go writer(ch, quit)
	go reader(ch, quit)
	<-time.After(d)
	quit <- true
}
