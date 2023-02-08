package cache

import (
	"errors"
	"time"

	"github.com/LebedevOleg/WB_test_task/internal/models"
)

var cache OrdersCache

type OrdersCache struct {
	Orders     []models.JsonOrder
	expiretion time.Duration
	timeCheck  time.Time
}

func GetAllOrders() []models.JsonOrder {
	return cache.Orders
}

func GetOrder(oid string) (models.JsonOrder, error) {
	for _, v := range cache.Orders {
		if v.Order_uid == oid {
			return v, nil
		}
	}
	return models.JsonOrder{}, errors.New("not Found")
}

func ConfigCache(expiretion time.Duration) {
	cache.expiretion = expiretion
}

func (oc *OrdersCache) SetExpiration(expiretion time.Duration) {
	oc.expiretion = expiretion
}

func UpdateCache(orders []models.JsonOrder) {
	cache.Orders = orders
	cache.timeCheck = time.Now()
}
func CheckUpdate() bool {
	return time.Since(cache.timeCheck) >= cache.expiretion
}
