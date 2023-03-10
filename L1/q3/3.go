package main

import (
	"fmt"
	"sync"
)

func sqrtSum() func(int, *sync.WaitGroup) int {
	res := 0
	return func(i int, wg *sync.WaitGroup) int {
		defer wg.Done()
		wg.Add(1)
		res = res + i*i
		return res
	}
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	//* используем замыкание для того чтобы расчитать сумму
	res := sqrtSum()
	for _, v := range arr {
		res(v, &wg)
	}
	//* альтернативный способ подсчета, по такому же принципу
	/* res := 0
	for _, v := range arr {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			res = res + v*v
		}(v)
	} */
	wg.Wait()
	fmt.Println(res(0, &wg))
	//fmt.Println(res)
}
