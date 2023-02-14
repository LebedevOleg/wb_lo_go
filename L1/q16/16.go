package main

import (
	"fmt"
	"sort"
)

func quickSort(arr []int, left, right int) { //* передаем срез наименьший индекс и наибольший
	if left >= right {
		return
	}
	split := partitional(arr, left, right) //* ищем индекс по которому делим срез на 2 части
	quickSort(arr, left, split)            //* запускаем рекурсию с левой частью среза (относительно найденной точки деления)
	quickSort(arr, split+1, right)         //* запускаем рекурсию с правой частью среза (относительно найденной точки деления)
	//* повторяем эту рекурсию пока наименьший индекс меньше наибольшего
}

func partitional(arr []int, left, right int) int {
	mid := arr[(left+right)/2] //* берем центральный элемент
	for left <= right {        //* пока наименьший индекс меньше или равен наибольшему
		//* увеличиваем наименьший индекс на 1 пока элемент под этим индексом меньше центрального
		for ; arr[left] < mid; left++ {
		}
		//* уменьшаем наибольший индекс на 1 пока элемент под этим индексом больше центрального
		for ; arr[right] > mid; right-- {
		}
		//* если наименьший индекс больше или равен наибольшему возвращаем наибольший индекс как делитель среза на 2 части
		if left >= right {
			return right
		}
		//* иначе меняем местами элемент с наименьшим индексом и с наибольшим
		arr[left], arr[right] = arr[right], arr[left]
	}
	return right
}

func main() {
	//* быстрая сортировка
	arr := []int{1, 20, 5, 7, 0, -2, 3, 4, 10, 2}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	//* Сравнение результата с встроенным методом
	arr = []int{1, 20, 5, 7, 0, -2, 3, 4, 10, 2}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)

}
