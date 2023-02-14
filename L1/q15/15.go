package main

import "fmt"

var justString string

func someFunc() {

	v := 1 << 10
	fmt.Println(v)
	//justString = v[:100]
}

func main() {
	someFunc()
}
