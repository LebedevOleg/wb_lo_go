package main

import (
	"fmt"
	"sync"
)

func chanel1(ch chan int, nextCh chan int) {
	x := <-ch
	nextCh <- x
}

func chanel2(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := <-ch
	result = result * result
	fmt.Println(result)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg sync.WaitGroup
	for _, v := range arr {
		wg.Add(1)
		go chanel1(ch1, ch2)
		go chanel2(ch2, &wg)
		ch1 <- v
	}
	wg.Wait()
}
