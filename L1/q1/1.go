package main

import (
	"fmt"
	"strconv"
)

type Human struct { //* Структура родитель
	Name string
	Age  int
	Sex  string
	Race string
}

func (h *Human) GetAge() int { //* метод структуры родителя
	return h.Age //! Возвращает ТОЛЬКО ВОЗРАСТ
}

type Action struct { //* Структура наследник
	Human
	//* Human Human -Аналогичная запись строки выше
}

func (a *Action) GetAge() string { //* метод структуры наследника

	return "человеку " + a.Name + " " + (strconv.Itoa(a.Age)) + " лет" //! Возвращает имя и возраст
}

func main() {
	var human = Human{"oleg", 10, "male", "afro"} // задаем значение Human
	var action = Action{human}                    // задаем значение Action

	fmt.Println("Функция родителя: ", human.GetAge())          //* При вызове функции получим возраст объекта human
	fmt.Println("Функция наследника: ", action.GetAge())       //! В данном случае отработает функция структуры Action, Поскольку она более приоритетна
	fmt.Println("Функция наследника: ", action.Human.GetAge()) //! Таким образом можно вызвать метод класса родителя
	/*
	* Если бы у структуры Action не было метода с одинаковыми названиями то родительский можно было бы вызвать следующей строчкой
	*fmt.Println("Функция наследника: ", action.GetAge())
	 */

	human.Age = 11 //! Изменяется только возраст объекта human, возраст action не поменяется

	fmt.Println("Функция родителя: ", human.GetAge())
	fmt.Println("Функция наследника: ", action.GetAge())
}
