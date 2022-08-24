Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
В выводе что-то похожее на последовательность от 1 до 8, но возможны некоторые перестановки в связи с рандомным таймером выполнения
1
2
4
3
6
5
8
7

Однако после этого выводится бесконечно 0. Хотя у нас закрыты оба канала. Они небуф. После их закрытия select считывает значения 0 и записывает в канал c
Хорошим решением проблемы было бы обнуление каналов.
default не будет работать, так как с канала приходит значение 0
Почему приходит значение 0 из закрытого канала? Потому что так работают закрытые каналы
val, ok := <-c
Придёт значение 0 и сигнал о том, что канал закрыт