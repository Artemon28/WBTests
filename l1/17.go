package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	array := make([]int, 20)
	for i := 0; i < 20; i++ {
		array[i] = rand.Intn(1000)
	}

	//бин поиск
	sort.Ints(array)
	fmt.Println(array)
	x := array[2]
	i := sort.Search(len(array), func(i int) bool { return x <= array[i] })

	fmt.Println(i)
}
