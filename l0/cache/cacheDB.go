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
)

func Recover(client postgresDataBase.Client) (map[string]cacheModel.Order, error) {
	newRepository := order.NewRepository(client)
	orders, err := newRepository.FindAll(context.TODO())
	if err != nil {
		return nil, err
	}
	for key, value := range orders {
		del, err := delivery.NewRepository(client).FindOne(context.TODO(), value.Del_phone)
		if err != nil {
			return nil, err
		}
		pay, err := payment.NewRepository(client).FindOne(context.TODO(), value.Order_uid)
		if err != nil {
			return nil, err
		}
		items, err := item.NewRepository(client).FindAll(context.TODO(), value.Order_uid)
		if err != nil {
			return nil, err
		}
		value.Delivery = del
		value.Payment = pay
		value.Items = items
		orders[key] = value
	}
	return orders, nil
}

func WriteToDB(client postgresDataBase.Client, newOrder cacheModel.Order) error {
	err := delivery.NewRepository(client).Create(context.TODO(), &newOrder.Delivery)
	if err != nil {
		return err
	}
	err = order.NewRepository(client).Create(context.TODO(), &newOrder)
	if err != nil {
		return err
	}
	err = payment.NewRepository(client).Create(context.TODO(), &newOrder.Payment)
	if err != nil {
		return err
	}
	for _, newItem := range newOrder.Items {
		err = item.NewRepository(client).Create(context.TODO(), &newItem)
		if err != nil {
			return err
		}
		err = items_in_order.NewRepository(client).Create(context.TODO(), newOrder.Order_uid, newItem.Chrt_id)
		if err != nil {
			return err
		}
	}
	return nil
}
