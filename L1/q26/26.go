package main

import (
	"fmt"
	"strings"
)

func main() {
	letters := make(map[rune]bool)
	var textLine string
	result := true
	fmt.Print("Введите строку: ")
	fmt.Scan(&textLine)
	for _, v := range strings.ToLower(textLine) {
		if _, ok := letters[v]; !ok {
			letters[v] = true
			continue
		}
		result = false
		break
	}
	fmt.Println(textLine+" - ", result)
}
