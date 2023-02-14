package main

import (
	"fmt"
	"strings"
)

func main() {
	letters := make(map[rune]bool) //* map для записи букв. ключ - буква значение - булево переменная, true есть
	var textLine string
	result := true //* булева переменная результата
	fmt.Print("Введите строку: ")
	fmt.Scan(&textLine)
	for _, v := range strings.ToLower(textLine) { //* проходим по строке приведенную к нижнему регистру
		if _, ok := letters[v]; !ok { //* если буква еще не записано, то записываем ее
			letters[v] = true
			continue
		}
		//* иначе буква уже есть, а значит строка с повтором символа
		result = false
		break
	}
	fmt.Println(textLine+" - ", result)
}
