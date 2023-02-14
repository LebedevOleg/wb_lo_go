package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func worker(ch chan int, id string) {
	for { //* запускаем бесконечный цикл, чтобы горутина не завершилась после первого чтения
		//* проверяем канал на активность
		if data, ok := <-ch; ok {
			//* если канал открыт, то выводим имя горутины и данные из канала
			fmt.Print("W"+id+": ", data, ", ")
			continue
		}
		//* если канал закрыт то останавливаем горутину
		return
	}
}

func main() {
	var workerCount int
	max := 99999
	fmt.Print("Введите коллиество воркеров")
	fmt.Scan(&workerCount)
	ch := make(chan int)         //* создаем канал для передачи данных
	c := make(chan os.Signal, 1) //* создаем переменную которая будет считывать сигналы из системы
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	for i := 0; i < workerCount; i++ {
		//* запускаем выбранное колличество рабочих горутин
		go worker(ch, strconv.Itoa(i))
	}
	for { //* запускаем бесконечный цикл
		rand.Seed(time.Now().UnixNano()) //* задаем зерно для рандома
		ch <- rand.Intn(max)             //* пишем случайно сгенерированное число в канал
		select {
		//* ждем сигнала что программу надо остановить (ctrl+c)
		case <-c:
			close(ch)
			os.Exit(1)
		//* если сигнала нет, то ждем 30 милисекунд и продолжаем
		default:
			time.Sleep(30 * time.Millisecond)
			continue
		}
	}
}
