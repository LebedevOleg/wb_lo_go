package main

import (
	"fmt"
	"sync"
	"time"
)

func sqrt(value int) {
	value = value * value
	fmt.Println(value)
}

func gorutine(arr []int, wg sync.WaitGroup) {
	for _, v := range arr {
		wg.Add(1)
		go sqrt(v)
	}
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	timer := time.Now()
	gorutine(arr, wg)
	wg.Wait()
	fmt.Print(time.Since(timer).Milliseconds())
}
