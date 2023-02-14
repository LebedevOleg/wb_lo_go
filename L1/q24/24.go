package main

import (
	"fmt"
	"math"
)

// ! название структуры и переменных внутри с маленькой буквы
// ! это нужно для того, чтобы при экспорте структуры нельзя было получить прямого доступа к ней и ее переменным
type point struct { //* структура Point{x, y}
	x float64
	y float64
}

// ! чтобы конструктор полностью себя реализовал, структура и его конструктор должны находиться в другом пакете
// ! Название ф-ции с большой буквы, видно при экспорте
func NewPoint(x float64, y float64) point { //* конструктор новой точки
	return point{x, y}
}

func CalculateDistance(p1 point, p2 point) float64 { //* считает расстояние между двумя точками
	return math.Sqrt(math.Pow((p2.x-p1.x), 2) + math.Pow((p2.y-p1.y), 2))
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(4, 4)
	fmt.Println(CalculateDistance(p1, p2))
}
