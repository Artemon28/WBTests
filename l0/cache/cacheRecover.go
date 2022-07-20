package cache

import (
	"L0/cache/delivery"
	"L0/cache/item"
	"L0/cache/items_in_order"
	"L0/cache/order"
	"L0/cache/payment"
	"L0/cacheModel"
	"L0/postgresDataBase"
	"context"
	"log"
)

//запускает восстановление, то есть, собирает все заказы
//из них достаёт все связанные доставки, сохраняет в этот заказ
//достаёт также оплату и присоединяет к заказу
//находит все итемы по id заказа, складывает в массив и также присоединяет к заказу
//всё складируется в мапе - наш кэш

func Recover(client postgresDataBase.Client) map[string]cacheModel.Order {
	newRepository := order.NewRepository(client)
	orders, _ := newRepository.FindAll(context.TODO())
	for _, value := range orders {
		del, _ := delivery.NewRepository(client).FindOne(context.TODO(), value.Del_phone)
		pay, _ := payment.NewRepository(client).FindOne(context.TODO(), value.Order_uid)
		items, _ := item.NewRepository(client).FindAll(context.TODO(), value.Order_uid)
		value.Delivery = del
		value.Payment = pay
		value.Items = items
	}
	return orders
}

func WriteToDB(client postgresDataBase.Client, newOrder cacheModel.Order) {
	newRepository := order.NewRepository(client)
	err := newRepository.Create(context.TODO(), &newOrder)
	if err != nil {
		log.Print("trouble")
		return
	}
	err = delivery.NewRepository(client).Create(context.TODO(), &newOrder.Delivery)
	if err != nil {
		log.Print("trouble")
		return
	}
	err = payment.NewRepository(client).Create(context.TODO(), &newOrder.Payment)
	if err != nil {
		log.Print("trouble")
		return
	}
	for _, newItem := range newOrder.Items {
		err1 := item.NewRepository(client).Create(context.TODO(), &newItem)
		err2 := items_in_order.NewRepository(client).Create(context.TODO(), newOrder.Order_uid, newItem.Chrt_id)
		if err1 != nil || err2 != nil {
			log.Print("trouble")
			return
		}
	}

}
