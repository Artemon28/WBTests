package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	//justString не копирует данные, а является указателем на v, то есть, если мы изменим v, то изменится и значение justString.
	//Однако строки это неизменяемые слайсы битов.
	justString = v[:100]
	fmt.Println(&justString)
	fmt.Println(&v)
	justString = strings.Replace(justString, justString[:10], "1234567890", -1)
	fmt.Println(&justString)
	fmt.Println(&v)
}

func createHugeString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := rand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func main() {
	someFunc()
}
