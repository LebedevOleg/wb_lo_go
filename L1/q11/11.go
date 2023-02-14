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
	//! в задаче делается акцент на множестве, значит слова внутри однго множества не повторяются
	arr1 := []string{"cat", "dog", "sun", "up"}          //* первое множество
	arr2 := []string{"cut", "down", "up", "left", "sun"} //*второе множество
	UnitedArr := make(map[string]int)                    //* map для двух множеств. Ключ - это слово, значение колличество повторений
	//! в данной задаче это ни на что не повлияет, но в теории за счет выделения памяти на моменте создания экономится время за счет отсутвия операции расширения
	resultArr := make([]string, 0, cap(func() []string { //* создаем срез, объем которого будет максимальная длинна из 2 срезов
		if len(arr1) > len(arr2) {
			return arr1
		}
		return arr2
	}()))
	//* заполняем map с помощью горутин используя для предотвращения состояния гонки
	var wg sync.WaitGroup
	var mute sync.RWMutex
	wg.Add(2)
	go fillMap(arr1, UnitedArr, &wg, &mute) //* добавляем в map данные из первого множества
	go fillMap(arr2, UnitedArr, &wg, &mute) //* добавляем в map данные из второго множества
	wg.Wait()
	for k, v := range UnitedArr { //* проходим по map и если значение в ключе больше одного то это слово есть в двух множествах
		if v > 1 {
			resultArr = append(resultArr, k)
		}
	}
	fmt.Println(resultArr)
}
