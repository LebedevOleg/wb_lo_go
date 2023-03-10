package main

import (
	"fmt"
	"sort"
)

func arraySort(arr []int) { //* сортировка встроенной в библиотеку sort методом (быстрая сортировка)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

func binarySerch(arr []int, value int) (int, int, bool) {
	left, right := 0, len(arr)-1 //* задаем наименьшее и наибольшие значения
	var mid int
	for left <= right { //* повторяем пока наименьшее значение меньше или равно наибольшему
		mid = left + (right-left)/2 //* вычисляем центарльный индекс

		if value < arr[mid] { //* если число которое мы ищем меньше центарльного
			right = mid - 1 //* то наибольший индекс сдвигается на 1 от центрального и переходим дальше
			continue
		} else if value > arr[mid] { //* если число которое мы ищем больше центарльного
			left = mid + 1 //* то наименьший индекс сдвигается на 1 от центрального и переходим дальше
			continue
		}
		return arr[mid], mid, true //* соответсвенно если мы не продолжили то мы нашли нужное число и возвращаем его, его индекс в срезе и булево состояние успеха поиска
	}
	return -1, -1, false //* если мы ничего не вернули, значит числа что мы ищем нет. возвращается -1 значение и индекс и false как состояние не успеха

}

func main() {
	arr := []int{1, 2, 4, 3, 5, 6, 7, 8, 9, 10} //* срез для поиска
	var serchNum int
	fmt.Print("Введите число для поиска: ")
	fmt.Scan(&serchNum)
	arraySort(arr)                                    //* поскольку бинарный поиск выполняется по отсортированному срез, то необходимо отсортировать его
	if _, ind, ok := binarySerch(arr, serchNum); ok { //* проверяем состояние успеха если true то выводится значение
		fmt.Println("Число находится на позиции:", ind)
		return
	}
	fmt.Println("Числа нет в массиве")

}
