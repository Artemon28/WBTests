package cache

import (
	"L0/cacheModel"
	"encoding/json"
)

func ConvertFromJson(data []byte) (cacheModel.Order, error) {
	var orderFromJson cacheModel.Order
	err := json.Unmarshal(data, &orderFromJson)
	if err != nil {
		return cacheModel.Order{}, err
	}
	err = cacheModel.ModelCheck(orderFromJson)
	if err != nil {
		return cacheModel.Order{}, err
	}
	orderFromJson.Del_phone = orderFromJson.Delivery.Phone
	return orderFromJson, nil
}
