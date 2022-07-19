package main

import (
	"L0/order"
	"L0/postgresDataBase"
)

func main() {
	//Read from canal
	//Add this to struct and check if everything is ok
	//write to DB and Map

	var orderCache map[string]order.Order
	newOrder := order.Order{}
	orderCache[newOrder.GetID()] = newOrder
	postgresDataBase.WriteToDB(newOrder)
}
