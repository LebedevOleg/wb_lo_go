package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b, result big.Int //* создаем переменные
	var aNum, bNum, operator string
	//! вводим цифру оператор и цифру через пробел
	fmt.Print(`	Программа перемножает, делит, складывает, вычитает две числовых переменных
	Поддерживаемые знаки (+, *, -, /)
	Пример запроса: 10 + 10
	Введите запрос:`)
	fmt.Scan(&aNum)
	fmt.Scan(&operator)
	fmt.Scan(&bNum)
	a.SetString(aNum, 10)
	b.SetString(bNum, 10)

	switch operator { //* в зависимости от оператора выбираем действие
	case "+":
		result.Add(&a, &b)
	case "*":
		result.Mul(&a, &b)
	case "-":
		result.Sub(&a, &b)
	case "/": //* важно отметить что при делении на ноль типа float мы не получим ошибку а получим +бесконечность
		if b.BitLen() == 0 { //* если переменные были бы int то необходимо делать проверку на деление на 0
			fmt.Println("АЙ АЙ АЙ не надо делить на 0")
			return
		}
		result.Div(&a, &b)
	}
	fmt.Println("Результат: ", result.String())
}
