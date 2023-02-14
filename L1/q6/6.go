package main

import (
	"context"
	"fmt"
)

func chanGorotine(ch chan int) { //* Пример остановки горутины с помощью проверки канала
	fmt.Println("start gortine with channel end")
	ch <- 1                 //! необходим для того чтобы увидеть выводы
	if _, ok := <-ch; !ok { //* ждет ответа от канала, и если он закрыт то завершает горутину
		fmt.Println("stop gortine with channel end")
		return
	}
}

func selectGorutine(ch chan struct{}) { //* Закрытие горутины через блок select
	fmt.Println("start gortine with select end")
	//! в данном примере select только и нужен чтобы закрыть канал, но так же он может поддерживать дополнительные кейсы(как в задаче 5)
	select {
	case <-ch: //* ожидаем данных от канала и при получении завершаем горутину
		fmt.Println("stop gortine with select end")
		return
	}
}

func contextGorutine(ctx context.Context) { //* закрытие горутины с использованием контекста
	fmt.Println("start gortine with context end")
	select {
	case <-ctx.Done(): //* если контекст закрыт, то завершаем горутину
		fmt.Println("stop gortine with context end")
		return
	}
}

func main() {
	ch := make(chan int)
	go chanGorotine(ch)
	<-ch
	close(ch)
	//! для синхранизации основного потока и горутины используется ожидание ввода
	//! но можно использовать и waitgroup как в примере(3)
	fmt.Scanln()

	doneCH := make(chan struct{})
	go selectGorutine(doneCH)
	doneCH <- struct{}{}
	fmt.Scanln()

	ctx, cancel := context.WithCancel(context.Background())
	go contextGorutine(ctx)
	cancel()
	fmt.Scanln()
}
