package nats

import (
	"github.com/LebedevOleg/WB_test_task/internal/models"
	"github.com/brianvoe/gofakeit/v6"
)

func CreateRandomOrder() models.JsonOrder {
	var o models.JsonOrder

	gofakeit.Seed(0)
	gofakeit.Struct(&o)

	return o
}
