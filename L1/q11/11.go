package main

import (
	"fmt"
	"sync"
)

func fillMap(array []string, m map[string]int, wg *sync.WaitGroup, mute *sync.RWMutex) {
	defer wg.Done()

	for _, value := range array {
		mute.RLock()
		if _, ok := m[value]; ok {
			mute.RUnlock()
			mute.Lock()
			m[value]++
			mute.Unlock()
			continue
		}
		mute.RUnlock()
		mute.Lock()
		m[value] = 1
		mute.Unlock()

	}
}

func main() {
	arr1 := []string{"cat", "dog", "sun", "up"}
	arr2 := []string{"cut", "down", "up", "left", "sun"}
	UnitedArr := make(map[string]int)
	resultArr := make([]string, 0, cap(func() []string {
		if len(arr1) > len(arr2) {
			return arr1
		}
		return arr2
	}()))

	var wg sync.WaitGroup
	var mute sync.RWMutex
	wg.Add(2)
	go fillMap(arr1, UnitedArr, &wg, &mute)
	go fillMap(arr2, UnitedArr, &wg, &mute)
	wg.Wait()
	for k, v := range UnitedArr {
		if v > 1 {
			resultArr = append(resultArr, k)
		}
	}
	fmt.Println(resultArr)
}
