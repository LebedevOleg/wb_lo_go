package main

//! not work
import "fmt"

func main() {
	arr := []int{1, 20, 5, 7, 0, -2, 3, 4, 10, 2}
	fmt.Println(arr, len(arr))
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	arr = []int{1, 20, 5, 7, 0, -2, 3, 4, 10, 2}
	arr = qwik_sort(arr, 0, len(arr)-1)
	fmt.Println(arr)

}

func quickSort(arr []int, left int, right int) {
	l, r := left, right

	center := arr[left+right/2]

	for l < r {
		for ; arr[r] > center; r-- {
		}

		for ; arr[l] < center; l-- {
		}

		if l <= r {
			arr[r], arr[l] = arr[l], arr[r]
			l++
			r--
		}
	}

	if r > left {
		quickSort(arr, left, r)
	}
	if l > right {
		quickSort(arr, l, right)
	}

}

func qwik_sort(lst []int, left int, right int) []int {

	//Создаем копии пришедших переменных, с которыми будем манипулировать в дальнейшем.
	l := left
	r := right

	//Вычисляем 'центр', на который будем опираться. Берем значение ~центральной ячейки массива.
	center := lst[(left+right)/2]

	fmt.Println(l, r, (left+right)/2)

	//Цикл, начинающий саму сортировку
	for l <= r {

		//Ищем значения больше 'центра'
		for lst[r] > center {
			r--
		}

		//Ищем значения меньше 'центра'
		for lst[l] < center {
			l++
		}

		//После прохода циклов проверяем счетчики циклов
		if l <= r {
			//И если условие true, то меняем ячейки друг с другом.
			lst[r], lst[l] = lst[l], lst[r]
			l++
			r--
		}
	}

	if r > left {
		qwik_sort(lst, left, r)
	}

	if l > right {
		qwik_sort(lst, l, right)
	}

	return lst
}
