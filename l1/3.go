package main

import (
	"fmt"
)

func squareWorker(values <-chan int, results chan<- int) {
	v := <-values
	results <- v * v
}

func main() {
	array := []int{2, 4, 6, 8, 10, 11, 12, 13}
	values := make(chan int)
	results := make(chan int)
	for i := 1; i <= len(array); i++ {
		go squareWorker(values, results)
	}
	for i := 1; i <= len(array); i++ {
		values <- array[i-1]
	}
	close(values)

	sum := 0
	for i := 1; i <= len(array); i++ {
		sum += <-results
	}
	fmt.Println(sum)
}
