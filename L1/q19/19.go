package main

import "fmt"

func main() {
	var textLine string
	fmt.Print("Введите строку: ")
	fmt.Scan(&textLine)
	result := ""
	for _, v := range textLine {
		result = string(v) + result
	}
	fmt.Print(result)
}
