package main

import (
	"fmt"
	"time"
)

// по идее это не правльное решение функции sleep. так как при этом нагружен процессор.
// функция sleep запускает гоурутину таймер, которая возобновляет работу после истечения времени
func Sleep(duration time.Duration) {
	startTime := time.Now()
	for {
		if time.Now().Sub(startTime) >= duration {
			return
		}
	}
}

//Этот вариант странный, так как использует встроенную функцию After
func Sleep2(duration time.Duration) {
	<-time.After(duration)
	return
}

//А это другой встроенный вариант функции Sleep от time, но суть работы немного другая
func Sleep3(duration time.Duration) {
	t := time.NewTicker(duration)
	<-t.C
	return
}

func cpuLoader() {
	x := 2
	for {
		x = x * x
		if x > 3000000000 {
			x = 0
		}
	}
}

func main() {
	//for i := 1; i <= 30; i++ {
	//	go cpuLoader()
	//}
	var sum time.Duration
	start := time.Now()
	for i := 1; i <= 100; i++ {
		Sleep(time.Second)
		sum += time.Now().Sub(start)
		start = time.Now()
	}
	fmt.Println(sum / 100)
	//1.0003907s
	//with cpu Loader 1.016481222s

	sum = 0
	start = time.Now()
	for i := 1; i <= 100; i++ {
		Sleep2(time.Second)
		sum += time.Now().Sub(start)
		start = time.Now()
	}
	fmt.Println(sum / 100)
	//1.010525979s
	//with cpu Loader 1.033137331s

	sum = 0
	start = time.Now()
	for i := 1; i <= 100; i++ {
		Sleep3(time.Second)
		sum += time.Now().Sub(start)
		start = time.Now()
	}
	fmt.Println(sum / 100)
	//1.009247465s
	///with cpu Loader 1.026403968s

	sum = 0
	start = time.Now()
	for i := 1; i <= 100; i++ {
		time.Sleep(time.Second)
		sum += time.Now().Sub(start)
		start = time.Now()
	}
	fmt.Println(sum / 100)
	//1.010382481s
	//with cpu Loader 1.027224259s
}
