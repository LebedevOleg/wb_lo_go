package main

import (
	"fmt"
	"os"
	"time"
)

func Reader(ch chan int) {
	for {
		select {
		case number := <-ch:
			fmt.Println(number)
		}
	}
}

func main() {
	var inputTime int

	for n, err := fmt.Scanln(&inputTime); err != nil && n != 1; {
		fmt.Println("Введите одно число")
	}
	timer := time.After(time.Second * time.Duration(inputTime))
	fmt.Println(time.Second * time.Duration(inputTime))
	ch := make(chan int)
	num := 0
	go Reader(ch)
	for {
		select {
		case <-timer:
			fmt.Println("Программа завершена")
			close(ch)
			os.Exit(0)
		case ch <- num:
			num++
			continue
		}
	}
}
