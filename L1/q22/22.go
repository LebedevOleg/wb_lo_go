package main

import "fmt"

func main() {
	var a, b, result int64
	var operator string
	fmt.Print(`	Программа перемножает, делит, складывает, вычитает две числовых переменных
	Поддерживаемые знаки (+, *, -, /)
	Пример запроса: 10 + 10
	Введите запрос:`)
	fmt.Scan(&a)
	fmt.Scan(&operator)
	fmt.Scan(&b)
	fmt.Println(a, operator, b)

	switch operator {
	case "+":
		result = a + b
	case "*":
		result = a * b
	case "-":
		result = a - b
	case "/":
		if b == 0 {
			fmt.Println("АЙ АЙ АЙ не надо делить на 0")
			return
		}
		result = a / b
	}
	fmt.Println("Результат: ", result)
}
