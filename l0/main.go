package main

import (
	"L0/cache"
	"L0/cacheModel"
	"L0/postgresDataBase"
	"context"
	"github.com/nats-io/stan.go"
	"time"
)

func main() {
	connectionConfig := postgresDataBase.ConnectionConfig{Username: "postgres", Password: "UJBNVJ", Host: "localhost", Port: "5432", Database: "postgres"}
	connection, _ := postgresDataBase.NewClient(context.TODO(), connectionConfig)
	//
	////recover before start
	//
	orderCache := make(map[string]cacheModel.Order)
	//orderCache = cache.Recover(connection)
	//if orderCache == nil {
	//	os.Exit(1)
	//}
	//
	//Read order from canal

	var order cacheModel.Order
	sc, _ := stan.Connect("prod", "pub-test2")
	defer sc.Close()

	sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
		//fmt.Printf("Received a message: %s\n", string(m.Data))
		order, _ = cache.ConvertFromJson(m.Data)
		orderCache[order.Order_uid] = order
		cache.WriteToDB(connection, order)
	})
	time.Sleep(time.Second * 10)
	defer sub.Close()

}
