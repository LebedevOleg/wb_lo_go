package main

import (
	"fmt"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} //* последовательность температуры
	resultMap := make(map[int][]float64)                                //* map для хранения результата. хранит срез температур
	for _, v := range arr {
		key := int(v) - int(v)%10         //* вычисляем ключ для записи(температуры с шагом 10, в которую входит проверяемое число)
		if _, ok := resultMap[key]; !ok { //* если такого ключа еще нет то мы создаем его и записываем в него новый срез
			resultMap[key] = []float64{v}
			continue
		}
		resultMap[key] = append(resultMap[key], v) //* если такой ключ есть, то добавляем в срез новый элемент
	}
	fmt.Println(resultMap)
}
