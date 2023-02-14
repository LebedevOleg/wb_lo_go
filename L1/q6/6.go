package main

import (
	"context"
	"fmt"
)

func chanGorotine(ch chan int) {
	fmt.Println("start gortine with channel end")
	ch <- 1
	if _, ok := <-ch; !ok {
		fmt.Println("stop gortine with channel end")
		return
	}
	fmt.Println("end gorutine with channel end")
}

func selectGorutine(ch chan struct{}) {
	fmt.Println("start gortine with select end")
	select {
	case <-ch:
		fmt.Println("stop gortine with select end")
		return
	}
}

func contextGorutine(ctx context.Context) {
	fmt.Println("start gortine with context end")
	select {
	case <-ctx.Done():
		fmt.Println("stop gortine with context end")
		return
	}
}

func main() {
	ch := make(chan int)
	go chanGorotine(ch)
	<-ch
	close(ch)
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
