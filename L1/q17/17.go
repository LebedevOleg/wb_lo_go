package main

import (
	"fmt"
	"sort"
)

func arraySort(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

func binarySerch(arr []int, value int) (int, int, bool) {
	left, right := 0, len(arr)-1
	var mid int
	for left <= right {
		mid = left + (right-left)/2

		if value < arr[mid] {
			right = mid - 1
			continue
		} else if value > arr[mid] {
			left = mid + 1
			continue
		}
		return arr[mid], mid, true
	}
	return -1, -1, false

}

func main() {
	arr := []int{1, 2, 4, 3, 5, 6, 7, 8, 9, 10}
	var serchNum int
	fmt.Print("Введите число для поиска: ")
	fmt.Scan(&serchNum)
	arraySort(arr)
	if _, ind, ok := binarySerch(arr, serchNum); ok {
		fmt.Println("Число находится на позиции:", ind)
		return
	}
	fmt.Println("Числа нет в массиве")

}
