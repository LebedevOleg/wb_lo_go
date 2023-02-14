package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	var index int
	fmt.Print("Введите индекс на удаление: ")
	fmt.Scan(&index)
	if index < 0 || index >= len(arr) { //* проверяем введенный индекс на принадлежность к границам среза
		fmt.Print("Индекс не попадает в границы среза")
		return
	}
	var result []int //* результирующий срез
	for i, _ := range arr {
		if i == index { //* когда находим необходимый индекс
			result = append(result, arr[:i]...)   //* добавляем в результирующий срез значения до индекса
			result = append(result, arr[i+1:]...) //* и после его
			break
		}
	}
	fmt.Print(result)
}
