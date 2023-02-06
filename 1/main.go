package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

var cache OrdersCache

func main() {
	app := fiber.New()
	cache.SetExpiration(time.Minute / 2)
	ConnectDB("host=localhost user=postgres password=0000 dbname=WB_L0_orders port=5432 sslmode=disable")
	cache.SyncWithDB()
	app.Get("/orders", GetAllOrders)
	app.Post("/addorder", CreateOrder)

	app.Listen(":3000")
}

/* 	route := gin.Default()

route.GET("/", func(context *gin.Context) {
	getTestJson(context)
})
route.GET("/orders", GetAllOrders)
route.POST("/addorder", CreateOrder)

ConnectDB("host=localhost user=postgres password=0000 dbname=WB_L0_orders port=5432 sslmode=disable")

route.Run() */
