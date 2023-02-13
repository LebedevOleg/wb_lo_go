package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var text string
	textScanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите строку: ")
	textScanner.Scan()
	text = textScanner.Text()
	text = strings.TrimSpace(text) + " "
	textMap := make(map[int]string)
	word := ""
	resultText := ""
	wordIndex := 0
	for _, v := range text {
		if v != ' ' {
			word = word + string(v)
			continue
		}
		textMap[wordIndex] = word
		word = ""
		wordIndex++
	}

	for ; wordIndex >= 0; wordIndex-- {
		resultText = resultText + " " + textMap[wordIndex]
	}
	fmt.Println(resultText)

}
