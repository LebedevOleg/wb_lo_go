package main

import (
	"fmt"
	"sync"
)

func sqrt(value int, wg *sync.WaitGroup) {
	defer wg.Done() //* используем defer, чтобы выполнить действие после выполнения ф-ции
	value = value * value
	fmt.Println(value)
}

func gorutine(arr []int) {
	//* переменная wg необходима для синхранизации горутин, чтобы была возможнасть увидеть выведенный результат
	var wg sync.WaitGroup
	for _, v := range arr {
		//* добавление задания для горутины
		wg.Add(1)
		//* запуск горутины
		go sqrt(v, &wg)
	}
	//* ожидание выполнения всех горутин
	wg.Wait()
}

func main() {
	arr := []int{2, 4, 6, 8, 10} //* срез с данными для которых расчитываем квадраты
	gorutine(arr)
}
