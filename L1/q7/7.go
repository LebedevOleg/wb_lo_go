package main

import (
	"fmt"
	"sync"
)

func main() {
	testMap := make(map[int]string)
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mutex.Lock()
			testMap[i] = "cool text number " + string(i)
			mutex.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(testMap)
}
