package main

import "fmt"

func main() {
	array := []string{"cat", "cat", "dog", "cat", "tree"}
	var set = make(map[string]struct{})
	for _, v := range array {
		set[v] = struct{}{}
	}
	fmt.Println(set)
}
