package cache

import (
	"L0/cacheModel"
	"encoding/json"
)

//Здесь необходимо оьработать модель, совпадение с ней или нет. Если есть лишние поля, то ошибка, если полей не хватает, то ошибка.
//Если это вообще не JSON то ошибка.
func ConvertFromJson(data []byte) (cacheModel.Order, error) {
	var orderFromJson cacheModel.Order
	err := json.Unmarshal(data, &orderFromJson)
	orderFromJson.Del_phone = orderFromJson.Delivery.Phone
	return orderFromJson, err
}
