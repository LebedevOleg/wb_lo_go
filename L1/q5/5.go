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
	fmt.Print("Введите время в секундах")

	for n, err := fmt.Scanln(&inputTime); err != nil && n != 1; { //* проверка на введенные данные
		fmt.Println("Введены не корректные данные")
	}
	timer := time.After(time.Second * time.Duration(inputTime)) //* устанавливаем таймер
	ch := make(chan int)
	num := 0
	go Reader(ch) //* запускаем горутину чтения данных из канала
	for {
		select {
		case <-timer: //* если таймер подошел к концу программа завершается
			fmt.Println("Программа завершена")
			close(ch) //* закрываем канал чтобы остановить горутину
			os.Exit(0)
		case ch <- num: //* если таймер еще не окончен, то записываем данные в канал
			num++ //* инкрементируем переменную
			continue
		}
	}
}
