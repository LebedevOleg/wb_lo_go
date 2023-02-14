package main

//! not work
import (
	"fmt"
	"sort"
)

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	split := partitional(arr, left, right)
	quickSort(arr, left, split)
	quickSort(arr, split+1, right)
}

func partitional(arr []int, left, right int) int {
	mid := arr[(left+right)/2]
	for left <= right {
		for ; arr[left] < mid; left++ {
		}

		for ; arr[right] > mid; right-- {
		}
		if left >= right {
			return right
		}

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
