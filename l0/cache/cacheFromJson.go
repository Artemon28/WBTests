package cache

import (
	"L0/cacheModel"
	"encoding/json"
)

func ConvertFromJson(data []byte) (cacheModel.Order, error) {
	var orderFromJson cacheModel.Order
	err := json.Unmarshal(data, &orderFromJson)
	orderFromJson.Del_phone = orderFromJson.Delivery.Phone
	return orderFromJson, err
}
