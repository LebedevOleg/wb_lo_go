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
	biteNumber := strconv.FormatInt(number, 2)
	fmt.Println("Введенное число в 2 системе: ", biteNumber)
	fmt.Print("Введите индекс бита на изменение: ")
	fmt.Scan(&index)
	fmt.Print("Введите бит, который хотите вставить: ")
	fmt.Scan(&bite)
	currindex := len(biteNumber) - index
	if currindex < 0 {
		addCount := -currindex + 1
		for ; addCount > 0; addCount-- {
			biteNumber = "0" + biteNumber
		}
		currindex = len(biteNumber) - index
	}
	changedNumber := fmt.Sprint(biteNumber[:currindex-1], bite, biteNumber[currindex:])
	newNumber, _ := strconv.ParseInt(changedNumber, 2, 0)
	fmt.Println("Измененное число в 2 системе: ", changedNumber)
	fmt.Println("Измененное число в 10 системе: ", newNumber)

}
