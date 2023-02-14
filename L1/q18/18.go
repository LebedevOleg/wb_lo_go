package main

import (
	"fmt"
	"sync"
)

// * структура имеет значение int которое инкрементируется и mutex которое блокирует одновременную запись
type enumerator struct {
	count int
	mutex *sync.Mutex
}

func (c *enumerator) increment() {
	c.mutex.Lock()   //* блокируем одновременную запись
	c.count++        //* инкрементируем
	c.mutex.Unlock() //* разблокируем
}

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	enum := enumerator{0, &mutex}
	for i := 0; i < 10000; i++ { //* запускаем цикл на инкрементацию
		wg.Add(1)
		go func() {
			enum.increment()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(enum.count)
}
