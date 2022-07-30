package main

import (
	"fmt"
	"reflect"
)

type inter interface {
	read()
}

func main() {
	var x interface{}
	var ch = make(chan int)
	x = ch
	fmt.Println(reflect.TypeOf(x))
	x = 5
	fmt.Println(reflect.TypeOf(x))
	x = "some string"
	fmt.Println(reflect.TypeOf(x))
	x = true
	fmt.Println(reflect.TypeOf(x))
}
