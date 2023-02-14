package main

import "fmt"

func swap(a *int, b *int) (int, int) {
	return *b, *a
}

func main() {
	a, b := 1, 2 //* создаем 2 переменные
	fmt.Println(a, b)

	a, b = b, a //* меняем их местами
	fmt.Println(a, b)

	a, b = swap(&a, &b) //* меняем с помощью ф-ции
	fmt.Println(a, b)
}
