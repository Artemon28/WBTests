package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action
от родительской структуры Human (аналог наследования).
*/

type Human struct {
	name   string
	age    int
	weight float32
	height float32
}

func (human Human) GetAge() int {
	return human.age
}

func (human Human) GetWeight() float32 {
	return human.weight
}

func (human Human) GetIMT() float32 {
	return human.weight / (human.height * human.height)
}

type Action struct {
	Human
}

func main() {
	someAction := new(Action)
	someAction.Human = Human{"Max", 20, 75, 1.8}
	fmt.Println(someAction.GetAge())
	fmt.Println(someAction.GetWeight())
	fmt.Println(someAction.GetIMT())
}
