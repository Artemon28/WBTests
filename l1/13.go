package main

import "fmt"

func main() {
	a := 31
	b := 28
	fmt.Printf("a was: %d\nb was: %d\n", a, b)
	a, b = b, a
	fmt.Printf("a become: %d\nb become: %d", a, b)
}
