package main

import (
	"fmt"
)

func main() {
	slice := []string{"a", "a"}
	p := &slice
	fmt.Println(p)
	func(slice []string) {
		p := &slice
		fmt.Println(p)
		slice = append(slice, "a")
		fmt.Print(slice)
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)

	fmt.Print(slice)
}
