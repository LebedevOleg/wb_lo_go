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
	for {
		if data, ok := <-ch; ok {
			fmt.Print("W"+id+": ", data, ", ")
			continue
		}
		return
	}
}

func main() {
	var workerCount int
	max := 99999
	fmt.Print("Введите коллиество воркеров")
	fmt.Scan(&workerCount)
	ch := make(chan int)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	for i := 0; i < workerCount; i++ {
		go worker(ch, strconv.Itoa(i))
	}
	for {
		rand.Seed(time.Now().UnixNano())
		ch <- rand.Intn(max)
		select {
		case <-c:
			close(ch)
			os.Exit(1)
		default:
			time.Sleep(30 * time.Millisecond)
			continue
		}
	}
}
