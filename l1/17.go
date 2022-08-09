package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
)

func binSearch(array []int, l int, r int, key int) (int, error) {
	for l != r {
		mid := (l + r) / 2
		if key <= array[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if array[l] == key {
		return l, nil
	}
	return -1, errors.New("no this value")
}

func main() {
	array := make([]int, 0, 20)
	for i := 0; i < 10; i++ {
		array = append(array, rand.Intn(1000))
	}

	//бин поиск
	sort.Ints(array)
	fmt.Println(array)
	x := 81
	i, _ := binSearch(array, 0, len(array)-1, x)

	fmt.Println(i)
}
