package main

import (
	"fmt"
	"time"
)

func sleep(t time.Duration) { //* ф-ция сна
	<-time.After(t) //* ждет окончание таймера
}

func main() {
	var stop time.Duration = 5 * time.Second //* задаем паузу
	fmt.Println("время ", time.Now())
	sleep(stop) //* запускаем функцию сна
	fmt.Println("время ", time.Now())

}
