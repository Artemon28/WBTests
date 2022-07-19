package order

import (
	"time"
)

type Order struct {
	order_uid          string    `json:"order_uid"`
	track_number       string    `json:"track_number"`
	entry              string    `json:"entry"`
	locale             string    `json:"locale"`
	internal_signature string    `json:"internal_signature"`
	customer_id        string    `json:"customer_id"`
	delivery_service   string    `json:"delivery_service"`
	shardkey           int       `json:"shardkey"`
	sm_id              int       `json:"sm_id"`
	date_created       time.Time `json:"date_created"`
	oof_shard          string    `json:"oof_shard"`
	delivery           Delivery
	payment            Payment
	items              []Item
}

func (o *Order) GetID() string {
	return o.order_uid
}

type Delivery struct {
	name    string `json:"name"`
	phone   string `json:"phone"`
	zip     int64  `json:"zip"`
	city    string `json:"city"`
	address string `json:"address"`
	region  string `json:"region"`
	email   string `json:"email"`
}

type Payment struct {
	transaction   string `json:"transaction"`
	request_id    int    `json:"request_id"`
	currency      string `json:"currency"`
	provider      string `json:"provider"`
	amount        int    `json:"amount"`
	payment_dt    int64  `json:"payment_dt"`
	bank          string `json:"bank"`
	delivery_cost int    `json:"delivery_cost"`
	goods_total   int    `json:"goods_total"`
	custom_fee    int    `json:"custom_fee"`
}

type Item struct {
	chrt_id      int    `json:"chrt_id"`
	track_number string `json:"track_number"`
	price        int    `json:"price"`
	rid          string `json:"rid"`
	name         string `json:"name"`
	sale         int    `json:"sale"`
	size         string `json:"size"`
	total_price  int    `json:"total_price"`
	nm_id        int64  `json:"nm_id"`
	brand        string `json:"brand"`
	status       int    `json:"status"`
}
