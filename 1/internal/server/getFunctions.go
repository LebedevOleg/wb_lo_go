package server

import (
	"net/http"

	"github.com/LebedevOleg/WB_test_task/internal/cache"
	"github.com/LebedevOleg/WB_test_task/internal/database"
	"github.com/LebedevOleg/WB_test_task/internal/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllOrders(context *fiber.Ctx) error {
	if !cache.CheckUpdate() {
		context.Status(http.StatusOK).JSON(cache.GetAllOrders())
		return nil
	}

	var orders []models.Order
	var payment models.Payment
	var delivery models.Delivery
	var itemsOrders []models.Items_to_orders
	res := database.DB.Find(&orders)

	if res.Error != nil {
		context.Status(http.StatusBadRequest).JSON(res.Error.Error())
		//context.JSON(http.StatusBadRequest)
		return res.Error
	}

	jsonOrders := make([]models.JsonOrder, res.RowsAffected)
	for i, v := range orders {
		database.DB.Find(&payment, "transaction = ?", v.Payment_transctions)
		database.DB.Find(&delivery, "phone = ?", v.Delivery_phone)
		itemCount := database.DB.Find(&itemsOrders, "order_id = ?", v.Order_uid)
		items := make([]models.Item, itemCount.RowsAffected)
		for j, item := range itemsOrders {
			database.DB.Find(&items[j], "chrt_id = ? AND track_number = ?", item.Item_id, v.Track_number)
		}
		jsonOrders[i].ConvertToJsonOrder(orders[i], delivery, payment, items)
	}
	cache.UpdateCache(jsonOrders)
	context.Status(http.StatusOK).JSON(jsonOrders)
	return nil
	//	context.JSON(http.StatusOK, gin.H{"orders": jsonOrders})
}

func GetOrder(context *fiber.Ctx) error {
	orderID := context.Params("id")
	if !cache.CheckUpdate() {
		if res, err := cache.GetOrder(orderID); err == nil {
			context.Status(http.StatusOK)
			context.JSON(res)
			return nil
		}
	}
	defer database.SyncCache()
	var order models.Order
	var payment models.Payment
	var delivery models.Delivery
	var itemsOrders []models.Items_to_orders
	res := database.DB.Find(&order, "order_uid = ?", orderID)
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
	var jsonOrder models.JsonOrder

	database.DB.Find(&payment, "transaction = ?", order.Payment_transctions)
	database.DB.Find(&delivery, "phone = ?", order.Delivery_phone)
	itemCount := database.DB.Find(&itemsOrders, "order_id = ?", order.Order_uid)
	items := make([]models.Item, itemCount.RowsAffected)
	for j, item := range itemsOrders {
		database.DB.Find(&items[j], "chrt_id = ? AND track_number = ?", item.Item_id, order.Track_number)
	}
	jsonOrder.ConvertToJsonOrder(order, delivery, payment, items)
	context.Status(http.StatusOK)
	context.JSON(jsonOrder)
	return nil
}
