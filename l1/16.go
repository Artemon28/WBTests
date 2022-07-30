package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	array := make([]int, 100)
	for i := 0; i < 100; i++ {
		array[i] = rand.Intn(1000)
	}

	fmt.Println(array[:10])

	//сортировка
	sort.Ints(array)

	fmt.Println(array[:10])
}
