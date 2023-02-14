package main

import "fmt"

func main() {
	var textLine string
	fmt.Print("Введите строку: ")
	fmt.Scan(&textLine) //* вводим строку
	result := ""
	for _, v := range textLine { //* проходим посимвольно по строке и добавляем символ перед уже записанными данными
		result = string(v) + result
	}
	fmt.Print(result)
}
