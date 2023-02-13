package main

import "fmt"

type void interface{}

func main() {
	var number void = 3
	var text void = "text"
	var logical void = true
	var channel void = make(chan int)

	arr := []void{number, text, logical, channel}

	for _, v := range arr {
		switch v.(type) {
		case int:
			fmt.Printf("%v is a %[1]T\n", v)
		case string:
			fmt.Printf("%v is a %[1]T\n", v)
		case bool:
			fmt.Printf("%v is a %[1]T\n", v)
		case chan int:
			fmt.Printf("%v is a %[1]T\n", v)
		}
	}

}
