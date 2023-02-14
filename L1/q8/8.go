package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int64
	var index int
	var bite int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	biteNumber := strconv.FormatInt(number, 2) //* переводим число в двоичную систему в строковом виде
	fmt.Println("Введенное число в 2 системе: ", biteNumber)
	fmt.Print("Введите индекс бита на изменение: ")
	fmt.Scan(&index)
	fmt.Print("Введите бит, который хотите вставить: ")
	//! здесь нет проверки на введенное число, важно ввести 1 или 0
	fmt.Scan(&bite)
	currindex := len(biteNumber) - index //* высчитываем местоположение бита в строке
	if currindex < 0 {                   //* если значение отрицательное значит что индекс бита находится левее и нужно расширить строку до этого индекса
		addCount := -currindex + 1
		for ; addCount > 0; addCount-- {
			biteNumber = "0" + biteNumber
		}
		currindex = len(biteNumber) - index
	}
	changedNumber := fmt.Sprint(biteNumber[:currindex-1], bite, biteNumber[currindex:]) //* заменяем необходимый бит
	newNumber, _ := strconv.ParseInt(changedNumber, 2, 0)                               //* переводим обратно в число 10ичной системы
	fmt.Println("Измененное число в 2 системе: ", changedNumber)
	fmt.Println("Измененное число в 10 системе: ", newNumber)

}
