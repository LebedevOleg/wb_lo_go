package main

import (
	"fmt"
	"sync"
)

func chanel1(ch chan int, nextCh chan int) {
	x := <-ch   //* ожидаем получения числа
	nextCh <- x //* передаем его в следующий поток
}

func chanel2(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() //* помечаем задачу выполненой
	result := <-ch  //* читаем данные из потока
	result = result * result
	fmt.Println(result) //* выводит результат
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //* массив который пишется в канал 1
	var wg sync.WaitGroup                       //* необходим для синхронизации горутин с главным потокм
	for _, v := range arr {                     //* проходим по срезу и запускаем горутины на исполнения передовая нужные каналы
		wg.Add(1)
		go chanel1(ch1, ch2) //* передаем сюда канал в который пишем число и канал в который надо передать число
		go chanel2(ch2, &wg) //* передаем сюда какнал в которе будет записано число и waitgroup для синхронизации
		ch1 <- v             //* пишем число в первый поток
	}
	wg.Wait()
}
