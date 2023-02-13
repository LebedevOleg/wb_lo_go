package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	var index int
	fmt.Print("Введите индекс на удаление: ")
	fmt.Scan(&index)
	if index < 0 || index >= len(arr) {
		fmt.Print("Индекс не попадает в границы среза")
		return
	}
	var result []int
	for i, _ := range arr {
		if i == index {
			result = append(result, arr[:i]...)
			result = append(result, arr[i+1:]...)
			break
		}
	}
	fmt.Print(result)
}
