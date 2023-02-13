package main

import (
	"fmt"
	"sync"
)

type enumerator struct {
	count int
	mutex *sync.Mutex
}

func (c *enumerator) increment() {
	c.mutex.Lock()
	c.count++
	c.mutex.Unlock()
}

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	enum := enumerator{0, &mutex}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			enum.increment()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(enum.count)
}
