package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/LebedevOleg/WB_test_task/internal/nats"
	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster", "publisher", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		fmt.Println("ПОДКЛЮЧЕНИЕ НЕ УДАЛОСЬ")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Подключение прошло успешно")

	go SendNewOrder(sc)

	ctx, _ := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	select {
	case v := <-quit:
		log.Fatalf("signal not %v", v)
	case done := <-ctx.Done():
		log.Fatalf("it's done %v", done)
	}

	sc.Close()
}

func SendNewOrder(sc stan.Conn) {
	newRandomOrder := nats.CreateRandomOrder()
	if jsonOrder, err := json.Marshal(newRandomOrder); err != nil {
		fmt.Println(err.Error())
	} else {
		sc.Publish("publisher", jsonOrder)
	}
	fmt.Println("Сообщение отправлено")

}
