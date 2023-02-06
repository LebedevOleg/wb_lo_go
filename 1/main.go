package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/", func(context *gin.Context) {
		getTestJson(context)
	})
	route.GET("/orders", GetAllOrders)
	route.POST("/addorder", CreateOrder)

	ConnectDB("host=localhost user=postgres password=0000 dbname=WB_L0_orders port=5432 sslmode=disable")

	route.Run()
}

func GetAllOrders(context *gin.Context) {
	var orders []JsonOrder
	DB.Find(&orders)

	context.JSON(http.StatusOK, gin.H{"orders": orders})
}
