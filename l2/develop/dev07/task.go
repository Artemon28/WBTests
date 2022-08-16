package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	result := make(chan interface{})
	wg.Add(len(channels))

	//функция, которая читает всё, что записано в канал и записывает в канал с ответами
	output := func(sc <-chan interface{}) {
		for sqr := range sc {
			result <- sqr
		}
		wg.Done()
	}

	//запускаем эту функцию для каждого канала
	for _, optChan := range channels {
		go output(optChan)
	}

	//по завершению работы всех каналов закроем этот канал
	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(1*time.Minute),
		sig(1*time.Second),
		sig(50*time.Second),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))

}
