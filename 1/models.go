package main

type Payment struct {
	Transaction   string `json:"transaction" gorm:"primary_Key"`
	Request_id    string `json:"request_id"`
	Currency      string `json:"currency"`
	Provider      string `json:"provider"`
	Amount        int    `json:"amount"`
	Payment_dt    int    `json:"payment_dt"`
	Bank          string `json:"bank"`
	Delivery_cost int    `json:"delivery_cost"`
	Goods_total   int    `json:"goods_total"`
	Custom_fee    int    `json:"custom_fee"`
}

type Item struct {
	Chrt_id      int    `json:"chrt_id" gorm:"primary_Key"`
	Track_number string `json:"track_number" gorm:"primary_Key"`
	Price        int    `json:"price"`
	Rid          string `json:"rid"`
	Name         string `json:"name"`
	Sale         int    `json:"sale"`
	Size         string `json:"size"`
	Total_price  int    `json:"total_price"`
	Nm_id        int    `json:"nm_id"`
	Brand        string `json:"brand"`
	Status       int    `json:"status"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone" gorm:"primary_Key"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type JsonOrder struct {
	Order_uid          string   `json:"order_uid" gorm:"primary_Key" type:"varchar"`
	Track_number       string   `json:"track_number" type:"varchar"`
	Entry              string   `json:"entry" type:"varchar"`
	Delivery           Delivery `json:"delivery"`
	Payment            Payment  `json:"payment"`
	Item               []Item   `json:"items"`
	Locale             string   `json:"locale"`
	Internal_signature string   `json:"internal_signature"`
	Customer_id        string   `json:"customer_id"`
	Delivery_service   string   `json:"delivery_service"`
	Shardkey           string   `json:"shardkey"`
	Sm_id              int      `json:"sm_id"`
	Date_created       string   `json:"date_created" type:"timestamp"`
	Oof_shard          string   `json:"oof_shard"`
}
type Order struct {
	Order_uid           string `json:"order_uid" gorm:"primary_Key" type:"varchar"`
	Track_number        string `json:"track_number" type:"varchar"`
	Entry               string `json:"entry" type:"varchar"`
	Delivery_phone      string `json:"phone" type:"varchar"`
	Payment_transctions string `json:"transaction" type:"varchar"`
	Locale              string `json:"locale" type:"varchar"`
	Internal_signature  string `json:"internal_signature" type:"varchar"`
	Customer_id         string `json:"customer_id" type:"varchar"`
	Delivery_service    string `json:"delivery_service" type:"varchar"`
	Shardkey            string `json:"shardkey" type:"varchar"`
	Sm_id               int    `json:"sm_id" type:"varchar"`
	Date_created        string `json:"date_created" type:"timestamp"`
	Oof_shard           string `json:"oof_shard" type:"varchar"`
}

type Items_to_orders struct {
	Order_id          string `json:"order_uid"`
	Item_id           int    `json:"chrt_id"`
	Item_track_number string `json:"track_number"`
}

func (jo *JsonOrder) ConvertToJsonOrder(o Order, d Delivery, p Payment, i []Item) {
	jo.Customer_id = o.Customer_id
	jo.Date_created = o.Date_created
	jo.Delivery = d
	jo.Delivery_service = o.Delivery_service
	jo.Track_number = o.Track_number
	jo.Entry = o.Entry
	jo.Payment = p
	jo.Internal_signature = o.Internal_signature
	jo.Item = i
	jo.Locale = o.Locale
	jo.Shardkey = o.Shardkey
	jo.Sm_id = o.Sm_id
	jo.Oof_shard = o.Oof_shard
	jo.Order_uid = o.Order_uid
}
