package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := 167772168888
	b := "99999999999999999999"

	newA := big.NewInt(int64(a))
	newB := new(big.Int)
	newB.SetString(b, 10)

	result := new(big.Int)

	fmt.Printf("A is %d\nB is %d\n", newA, newB)
	//Sum
	result.Add(newA, newB)
	fmt.Printf("a + b is equals to: %d\n", result)

	//sub
	result.Sub(newB, newA)
	fmt.Printf("b - a is equals to: %d\n", result)

	//Multi
	result.Mul(newA, newB)
	fmt.Printf("b * a is equals to: %d\n", result)

	//Div
	floatB := new(big.Float)
	floatB.SetString(b)
	floatResult := new(big.Float).Quo(floatB, big.NewFloat(float64(a)))
	fmt.Printf("b / a is equals to: %f\n", floatResult)
}
