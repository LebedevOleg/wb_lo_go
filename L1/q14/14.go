package main

import "fmt"

type void interface{} //* пустой интерфейс для того чтобы любая перменная могла быть записано под ним

func main() {
	//* создаем набор переменных разного типа
	var number void = 3
	var text void = "text"
	var logical void = true
	var channel void = make(chan int)

	arr := []void{number, text, logical, channel} //* создаем массив интерфейса со всеми созданными переменными

	for _, v := range arr { //* проходим по массиву
		switch v.(type) { //* проверяем интерфейс на тип и выводим в консоль
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
