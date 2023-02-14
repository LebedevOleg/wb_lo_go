package main

import "fmt"

//* адаптер для перевода температуры из фаренгейта в цельсию

type SensorAdapter interface { //* интерфейс адаптер для которого подходит класс CelciusSensor
	GetCelTempData() float32
}

func newAdapter(temp float32) SensorAdapter { //* создание объекта интерфейса адаптера
	return CelciusSensor{&FahrenheitSensor{temp}}
}

type FahrenheitSensor struct { //* класс который нужно адаптировать
	temp float32
}

func (f *FahrenheitSensor) GetFahTempData() float32 { //* функция класса, который нужно адаптировать
	return f.temp //* возвращает градусы по Фаренгейту
}

type CelciusSensor struct { //* класс для которого адаптируем
	*FahrenheitSensor
}

func (c CelciusSensor) GetCelTempData() float32 { //* функция класса для которого адаптируем
	return (c.GetFahTempData() - 32) * 5 / 9 //* возвращаем градусы по цельсию
}

func main() {
	adapter := newAdapter(66.66)          //* Создаем объект интерфейса адаптера с температурой по фаренгейту
	fmt.Println(adapter.GetCelTempData()) //* выводится температура по цельсию
}
