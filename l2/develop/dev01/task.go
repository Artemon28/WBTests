package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

/*
=== Базовая задача ===
Создать программу печатающую точное время с использованием NTP -библиотеки. Инициализировать как go module.
Использовать библиотеку github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна проходить проверки go vet и golint.
*/

func main() {
	fmt.Println(getTime())
}

func getTime() time.Time {
	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal(err) //возвращает код ошибки 1 в OS и пишет в STDERR
	}
	return t
}
