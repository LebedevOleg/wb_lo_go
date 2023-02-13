package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) point {
	return point{x, y}
}
func CalculateDistance(p1 point, p2 point) float64 {
	return math.Sqrt(math.Pow((p2.x-p1.x), 2) + math.Pow((p2.y-p1.y), 2))
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(4, 4)
	fmt.Println(CalculateDistance(p1, p2))
}
