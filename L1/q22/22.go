package main

import "fmt"

func main() {
	var a, b, result float64 //* создаем переменные
	var operator string
	//! вводим цифру оператор и цифру через пробел
	fmt.Print(`	Программа перемножает, делит, складывает, вычитает две числовых переменных
	Поддерживаемые знаки (+, *, -, /)
	Пример запроса: 10 + 10
	Введите запрос:`)
	fmt.Scan(&a)
	fmt.Scan(&operator)
	fmt.Scan(&b)

	switch operator { //* в зависимости от оператора выбираем действие
	case "+":
		result = a + b
	case "*":
		result = a * b
	case "-":
		result = a - b
	case "/": //* важно отметить что при делении на ноль типа float мы не получим ошибку а получим +бесконечность
		/* if b == 0 { //* если переменные были бы int то необходимо делать проверку на деление на 0
			fmt.Println("АЙ АЙ АЙ не надо делить на 0")
			return
		} */
		result = a / b
	}
	fmt.Println("Результат: ", result)
}
