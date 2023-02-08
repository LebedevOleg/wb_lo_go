package nats

import (
	"encoding/json"
	"fmt"

	"github.com/LebedevOleg/WB_test_task/internal/database"
	"github.com/LebedevOleg/WB_test_task/internal/models"
	"github.com/nats-io/stan.go"
)

func GetMessage(m *stan.Msg) {
	var o models.JsonOrder

	if err := json.Unmarshal(m.Data, &o); err != nil {
		fmt.Println("Ошибка при десерилизации: ", err.Error())
		return
	}
	fmt.Println("Сообщние успешно серилизировано")
	fmt.Println(o.Date_created)
	m.Ack()

	err := database.AddNewOrder(o)
	if err != nil {
		fmt.Println("Ошибка при добавлении в БД \n", err.Error())
		return
	}
	fmt.Println("Сообщние успешно серилизировано")

	/////todo: добавление в БД и обновление кэша

}
