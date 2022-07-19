package main

import (
	"L0/cacheModel"
	"L0/postgresDataBase"
	"context"
)

func main() {
	//Read from canal
	//Add this to struct and check if everything is ok
	//write to DB and Map
	connectionConfig := postgresDataBase.ConnectionConfig{Username: "postgres", Password: "UJBNVJ", Host: "postgres", Port: "5432", Database: "postgres"}
	connection, _ := postgresDataBase.NewClient(context.TODO(), connectionConfig)

	var orderCache map[string]cacheModel.Order
	newOrder := cacheModel.Order{}
	orderCache[newOrder.Order_uid] = newOrder
	postgresDataBase.WriteToDB(connection, newOrder)
}
