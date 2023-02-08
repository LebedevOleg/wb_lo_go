package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/LebedevOleg/WB_test_task/internal/cache"
	"github.com/LebedevOleg/WB_test_task/internal/database"
	"github.com/LebedevOleg/WB_test_task/internal/nats"
	"github.com/LebedevOleg/WB_test_task/internal/server"
	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/stan.go"
)

func main() {
	//sc, err := stan.Connect("test-cluster", "subscribe", stan.NatsURL("nats://nats-streaming:5000"))
	sc, err := stan.Connect("test-cluster", "subscribe", stan.NatsURL("nats://localhost:4222"))
	cache.ConfigCache(time.Minute)
	if err != nil {
		fmt.Println("ПОДКЛЮЧЕНИЕ НЕ УДАЛОСЬ")
		fmt.Println(err)
		return
	}
	database.ConnectDB("host=localhost user=postgres password=0000 dbname=WB_L0_orders port=5432 sslmode=disable")

	sub, _ := sc.Subscribe("publisher", nats.GetMessage, stan.StartWithLastReceived(), stan.SetManualAckMode())

	app := fiber.New()

	//app.Use(cors.ConfigDefault)
	app.Static("/", "./internal/static/pages")
	app.Get("/orders", server.GetAllOrders)
	app.Get("/orders/:id", server.GetOrder)
	app.Listen(":3000")

	ctx, _ := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	select {
	case v := <-quit:
		log.Fatalf("signal not %v", v)
	case done := <-ctx.Done():
		log.Fatalf("it's done %v", done)
	}
	sub.Close()

}
