package main

import (
	"fmt"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	resultMap := make(map[int][]float64)
	for _, v := range arr {
		key := int(v) - int(v)%10
		if _, ok := resultMap[key]; !ok {
			resultMap[key] = []float64{v}
			continue
		}
		resultMap[key] = append(resultMap[key], v)
	}
	fmt.Println(resultMap)
}
