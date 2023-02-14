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
	textScanner.Scan()                  //* вводим строку с пробелами
	text = textScanner.Text()           //* сохраняем введенную строку
	text = strings.TrimSpace(text)      //* удаляем лишние пробеллы в начале и в конце
	textMap := make(map[int]string)     //* map для хранения слов, ключ - порядковый номер в введенном тексте, значение - слово
	textArr := strings.Split(text, " ") //* создаем срез из слов
	wordIndex := 0                      //* индекс слова
	resultText := ""                    //* результат
	for _, v := range textArr {         //* записываем слово и даем ему номер
		textMap[wordIndex] = v
		wordIndex++
	}
	//* поскольку map хранит ключи в неотсортированном виде есть шанс вывести слова не в обратном порядке
	for ; wordIndex >= 0; wordIndex-- { //* поэтому проходим по индексу уменьшая его пока не дойдем до -1
		resultText = resultText + " " + textMap[wordIndex]
	}

	fmt.Println(resultText)

}
