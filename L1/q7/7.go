package main

import (
	"fmt"
	"sync"
)

func main() {
	testMap := make(map[int]string)
	var mutex sync.Mutex //* для запрета одновременной записи в map используется mutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mutex.Lock()                                 //* блокируем доступ к map
			testMap[i] = "cool text number " + string(i) //* записываем данные в map
			mutex.Unlock()                               //* разблокируем
		}(i)
	}
	wg.Wait()
	fmt.Println(testMap)
}
