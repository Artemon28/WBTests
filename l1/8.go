package main

import (
	"fmt"
	"math"
)

func main() {
	var x int64
	x = 6805
	l := int(math.Log2(float64(x)))
	mask := int64(math.Pow(2, float64(l)))
	i := int64(5)
	mask = mask >> (i - 1)
	fmt.Printf("%b, int value is %d", x, x)
	fmt.Println()
	//Чтобы сделать i-ый бит 1
	fmt.Printf("%b, int value is %d", x|mask, x|mask)
	//Чтобы сделать i-ый бит 0
	fmt.Println()
	fmt.Printf("%b, int value is %d", x&^mask, x&^mask)
}
