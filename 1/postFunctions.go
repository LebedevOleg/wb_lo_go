package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(context *gin.Context) {
	var input JsonOrder
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	delivery := input.Delivery
	payment := input.Payment
	DB.Create(&delivery)
	DB.Create(&payment)

	order := Order{
		input.Order_uid, input.Track_number, input.Entry,
		input.Delivery.Phone, input.Payment.Transaction, input.Locale,
		input.Internal_signature, input.Customer_id, input.Delivery_service,
		input.Shardkey, input.Sm_id, input.Date_created, input.Oof_shard,
	}
	DB.Create(&order)

	for _, item := range input.Item {
		DB.Create(&item)
		itemtoorder := Items_to_orders{input.Order_uid, item.Chrt_id}
		DB.Create(&itemtoorder)
	}
	context.JSON(http.StatusOK, gin.H{"order": input})
}
