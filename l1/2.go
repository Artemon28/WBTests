package main

import (
	"fmt"
	"time"
)

func square(v int) {
	fmt.Println(v * v)
}

func main() {
	array := []int{2, 4, 6, 8, 10}
	for _, v := range array {
		go square(v)
	}
	time.Sleep(time.Second)
}
