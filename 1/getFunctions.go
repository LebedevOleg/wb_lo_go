package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetAllOrders(context *fiber.Ctx) error {
	if !cache.CheckUpdate() {
		context.Status(http.StatusOK).JSON(cache.Orders)
		return nil
	}

	var orders []Order
	var payment Payment
	var delivery Delivery
	var itemsOrders []Items_to_orders
	res := DB.Find(&orders)

	if res.Error != nil {
		context.Status(http.StatusBadRequest).JSON(res.Error.Error())
		//context.JSON(http.StatusBadRequest)
		return res.Error
	}

	jsonOrders := make([]JsonOrder, res.RowsAffected)
	for i, v := range orders {
		DB.Find(&payment, "transaction = ?", v.Payment_transctions)
		DB.Find(&delivery, "phone = ?", v.Delivery_phone)
		itemCount := DB.Find(&itemsOrders, "order_id = ?", v.Order_uid)
		items := make([]Item, itemCount.RowsAffected)
		for j, item := range itemsOrders {
			DB.Find(&items[j], "chrt_id = ? AND track_number = ?", item.Item_id, v.Track_number)
		}
		jsonOrders[i].ConvertToJsonOrder(orders[i], delivery, payment, items)
	}
	cache.UpdateCache(jsonOrders)
	context.Status(http.StatusOK).JSON(jsonOrders)
	return nil
	//	context.JSON(http.StatusOK, gin.H{"orders": jsonOrders})
}

func getOrder(context *fiber.Ctx) error {
	orderID := context.Params("id")
	if !cache.CheckUpdate() {
		for _, v := range cache.Orders {
			if v.Order_uid == orderID {
				context.Status(http.StatusOK)
				context.JSON(v)
				return nil
			}
		}
	}
	defer cache.SyncWithDB()
	var order Order
	var payment Payment
	var delivery Delivery
	var itemsOrders []Items_to_orders
	res := DB.Find(&order, "order_uid = ?", orderID)
	if res.Error != nil {
		context.Status(http.StatusBadRequest)
		context.JSON(res.Error.Error())
		return res.Error
	}
	if res.RowsAffected == 0 {
		context.Status(http.StatusBadRequest)
		context.JSON(`{"message": "order не найден"}`)
		return res.Error
	}
	var jsonOrder JsonOrder

	DB.Find(&payment, "transaction = ?", order.Payment_transctions)
	DB.Find(&delivery, "phone = ?", order.Delivery_phone)
	itemCount := DB.Find(&itemsOrders, "order_id = ?", order.Order_uid)
	items := make([]Item, itemCount.RowsAffected)
	for j, item := range itemsOrders {
		DB.Find(&items[j], "chrt_id = ? AND track_number = ?", item.Item_id, order.Track_number)
	}
	jsonOrder.ConvertToJsonOrder(order, delivery, payment, items)
	context.Status(http.StatusOK)
	context.JSON(jsonOrder)
	return nil
}
