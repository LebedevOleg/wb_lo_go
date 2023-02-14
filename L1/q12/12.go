package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	//! resultMap := make(map[string]struct{}) - можно инициализировать map так, поскольку нас не интерисует значение в ключе
	resultMap := make(map[string]bool) //* map для определения повторов в срезе. Ключ - слово, значение - булева переменная, показывающся было ли уже записано слово
	for _, v := range arr {
		if _, ok := resultMap[v]; !ok { //* если значение еще нет в map то создаем его
			resultMap[v] = true
		}
	}

	for k, _ := range resultMap { //* проходим по map и выводим все ключи
		fmt.Print(k + ", ")
	}
}
