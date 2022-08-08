package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func writeToStdin(strs chan<- string) {
	for {
		var str string
		fmt.Scan(&str)
		strs <- str
	}
}

//Создаём waitGroup горутин, они запускаются, работают. По нажатию Ctrl+C мы закрываем поток, и ждём,
//когда все горутины в группе закончат работу. таким образом мы дождёмся завершения работы горутин,
//время выполнения одной горутины в данной задаче не велико, что позволяет быстро завершить работу
func main() {
	var n int
	fmt.Println("Insert amount of workers")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err.Error())
	}

	values := make(chan string, n)
	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range values {
				fmt.Println(v)
			}
		}()

	}

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT)

	fmt.Println("Enter your strings here:")
	go writeToStdin(values)

	<-shutdownSignal
	close(values)
	wg.Wait()
}
