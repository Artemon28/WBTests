package main

import (
	"fmt"
	"math"
)

func main() {
	var x int64
	x = 6805
	//Найдём старший разряд нашего числа в двоичной системе
	l := int(math.Log2(float64(x)))
	//составим маску как сдвиг старшего числа на нужный нам разряд
	mask := int64(math.Pow(2, float64(l)))
	i := int64(5)
	mask = mask >> (i - 1)
	fmt.Printf("%b, int value is %d", x, x)
	fmt.Println()
	//Чтобы сделать i-ый бит 1 операция ИЛИ
	fmt.Printf("%b, int value is %d", x|mask, x|mask)
	//Чтобы сделать i-ый бит 0 операция сброс Бита или - И НЕ
	fmt.Println()
	fmt.Printf("%b, int value is %d", x&^mask, x&^mask)
}
