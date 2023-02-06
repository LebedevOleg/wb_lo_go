package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

const testPay = `{
  "order_uid": "b563feb7b2b84b6test",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "b563feb7b2b84b6test",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}`

func getTestJson(context *gin.Context) {
	order := JsonOrder{}
	if err := json.Unmarshal([]byte(testPay), &order); err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{"order": order})

}

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
