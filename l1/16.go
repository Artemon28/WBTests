package main

import (
	"fmt"
	"math/rand"
)

func quickSort(array []int, l int, r int) {
	if l < r {
		q := partition(array, l, r)
		quickSort(array, l, q)
		quickSort(array, q+1, r)
	}
}

func partition(array []int, l int, r int) int {
	pivot := array[(l+r)/2]
	i := l
	j := r
	for i <= j {
		for array[i] < pivot {
			i++
		}
		for array[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		array[i], array[j] = array[j], array[i]
		i++
		j--
	}
	return j

}

func main() {
	array := make([]int, 0, 100)
	for i := 0; i < 10; i++ {
		array = append(array, rand.Intn(1000))
	}

	fmt.Println(array[:10])

	quickSort(array, 0, len(array)-1)

	fmt.Println(array[:10])
}
