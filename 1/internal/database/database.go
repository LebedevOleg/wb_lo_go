package database

import (
	"errors"
	"fmt"

	"github.com/LebedevOleg/WB_test_task/internal/cache"
	"github.com/LebedevOleg/WB_test_task/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("no connect to DB")
	}
	DB = db
}

func AddNewOrder(jo models.JsonOrder) error {
	var order models.Order
	if res := DB.Find(&order, "order_uid = ?", jo.Order_uid); res.RowsAffected != 0 {
		fmt.Println("Этот объект уже есть в БД")
		return nil
	}

	var delivery models.Delivery
	var payment models.Payment
	if res := DB.Find(&delivery, "phone = ?", jo.Delivery.Phone); res.RowsAffected == 0 {
		delivery = jo.Delivery
		DB.Create(&delivery)
	}
	if res := DB.Find(&payment, "transaction = ?", jo.Payment.Transaction); res.RowsAffected == 0 {
		payment = jo.Payment
		DB.Create(&payment)
	}

	if res := DB.Find(&order, "order_uid = ?", jo.Order_uid); res.RowsAffected != 0 {
		err := errors.New("попытка добавить уже существующий заказ")
		return err
	}
	order = models.Order{
		jo.Order_uid, jo.Track_number, jo.Entry,
		jo.Delivery.Phone, jo.Payment.Transaction, jo.Locale,
		jo.Internal_signature, jo.Customer_id, jo.Delivery_service,
		jo.Shardkey, jo.Sm_id, jo.Date_created, jo.Oof_shard,
	}
	DB.Create(&order)

	for _, item := range jo.Item {
		DB.Create(&item)
		itemtoorder := models.Items_to_orders{jo.Order_uid, item.Chrt_id, item.Track_number}
		DB.Create(&itemtoorder)
	}
	SyncCache()
	return nil
}

func SyncCache() {
	var orders []models.Order
	var payment models.Payment
	var delivery models.Delivery
	var itemsOrders []models.Items_to_orders
	res := DB.Find(&orders)

	if res.Error != nil {
		return
	}

	jsonOrders := make([]models.JsonOrder, res.RowsAffected)
	for i, v := range orders {
		DB.Find(&payment, "transaction = ?", v.Payment_transctions)
		DB.Find(&delivery, "phone = ?", v.Delivery_phone)
		itemCount := DB.Find(&itemsOrders, "order_id = ?", v.Order_uid)
		items := make([]models.Item, itemCount.RowsAffected)
		for j, item := range itemsOrders {
			DB.Find(&items[j], "chrt_id = ? AND track_number = ?", item.Item_id, v.Track_number)
		}
		jsonOrders[i].ConvertToJsonOrder(orders[i], delivery, payment, items)
	}
	cache.UpdateCache(jsonOrders)
	fmt.Println("Успешное обновление Кэша")
}
