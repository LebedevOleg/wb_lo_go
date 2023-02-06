package main

import (
	"time"
)

type OrdersCache struct {
	Orders     []JsonOrder
	expiretion time.Duration
	timeCheck  time.Time
}

func (oc *OrdersCache) SetExpiration(expiretion time.Duration) {
	oc.expiretion = expiretion
}

func (oc *OrdersCache) UpdateCache(orders []JsonOrder) {
	oc.Orders = orders
	oc.timeCheck = time.Now()
}
func (oc *OrdersCache) CheckUpdate() bool {
	if time.Since(oc.timeCheck) >= oc.expiretion {
		return true
	}
	return false
}

func (oc *OrdersCache) SyncWithDB() {
	var orders []Order
	var payment Payment
	var delivery Delivery
	var itemsOrders []Items_to_orders
	res := DB.Find(&orders)

	if res.Error != nil {
		return
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
	oc.timeCheck = time.Now()
}
