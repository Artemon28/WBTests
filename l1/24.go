package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func (point *point) GetX() float64 {
	return point.x
}

func (point *point) GetY() float64 {
	return point.y
}

func NewPoint(x float64, y float64) *point {
	return &point{
		x: x,
		y: y,
	}
}

func distance(p1 point, p2 point) float64 {
	return math.Sqrt(math.Pow((p1.GetX()-p2.GetX()), 2) + math.Pow(p1.GetY()-p2.GetY(), 2))
}

func main() {
	p1 := NewPoint(2.0, 4.0)
	p2 := NewPoint(2.0, 8.0)
	fmt.Println(distance(*p1, *p2))
}
