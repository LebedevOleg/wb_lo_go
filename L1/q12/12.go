package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	resultMap := make(map[string]bool)
	for _, v := range arr {
		if _, ok := resultMap[v]; !ok {
			resultMap[v] = true
		}
	}

	for k, _ := range resultMap {
		fmt.Print(k + ", ")
	}
}

//! check one more
