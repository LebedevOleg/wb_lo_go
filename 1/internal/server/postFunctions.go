package server

/*
import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(context *fiber.Ctx) error {
	defer cache.SyncWithDB()

	var input JsonOrder
	//if err := context.ShouldBindJSON(&input); err != nil {
	if err := context.BodyParser(&input); err != nil {
		context.Status(http.StatusBadRequest).JSON(err.Error())
		context.SendStatus(http.StatusBadRequest)
		//context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	// delivery := input.Delivery
	//payment := input.Payment

	var delivery Delivery
	var payment Payment
	if res := DB.Find(&delivery, "phone = ?", input.Delivery.Phone); res.RowsAffected == 0 {
		delivery = input.Delivery
		DB.Create(&delivery)
	}
	if res := DB.Find(&payment, "transaction = ?", input.Payment.Transaction); res.RowsAffected == 0 {
		payment = input.Payment
		DB.Create(&payment)
	}

	// var order Order
	//order.ConvertToOrder(input)
	var order Order
	if res := DB.Find(&order, "order_uid = ?", input.Order_uid); res.RowsAffected != 0 {
		context.Status(http.StatusBadRequest).JSON("Попытка добавить уже существующий заказ")
		//context.JSON(http.StatusBadRequest, gin.H{"message": "Попытка добавить уже существующий заказ"})
		err := errors.New("попытка добавить уже существующий заказ")
		return err
	}
	order = Order{
		input.Order_uid, input.Track_number, input.Entry,
		input.Delivery.Phone, input.Payment.Transaction, input.Locale,
		input.Internal_signature, input.Customer_id, input.Delivery_service,
		input.Shardkey, input.Sm_id, input.Date_created, input.Oof_shard,
	}
	DB.Create(&order)

	for _, item := range input.Item {
		DB.Create(&item)
		itemtoorder := Items_to_orders{input.Order_uid, item.Chrt_id, item.Track_number}
		DB.Create(&itemtoorder)
	}

	context.Status(http.StatusOK).JSON(input)
	return nil
	//context.JSON(http.StatusOK, gin.H{"order": input})
}
*/
