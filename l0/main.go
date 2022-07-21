package main

import (
	"L0/cache"
	"L0/cacheModel"
	"L0/postgresDataBase"
	"context"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"net/http"
)

func main() {
	connectionConfig := postgresDataBase.ConnectionConfig{Username: "postgres", Password: "UJBNVJ", Host: "localhost", Port: "5432", Database: "postgres"}
	connection, _ := postgresDataBase.NewClient(context.TODO(), connectionConfig)

	orderCache := make(map[string]cacheModel.Order)
	orderCache = cache.Recover(connection)
	//if orderCache == nil {
	//	os.Exit(1)
	//}
	//
	//Read order from canal

	sc, _ := stan.Connect("prod", "pub-test2")
	defer sc.Close()

	sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
		//fmt.Printf("Received a message: %s\n", string(m.Data))
		newOrder, _ := cache.ConvertFromJson(m.Data)
		orderCache[newOrder.Order_uid] = newOrder
		cache.WriteToDB(connection, newOrder)
	})
	defer sub.Close()

	var order cacheModel.Order

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "L0/index.html")
	})

	http.HandleFunc("/order/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		switch r.Method {
		case http.MethodGet:
			orderJson, _ := json.Marshal(order)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(orderJson)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)
	//nats-streaming-server -cid prod
}
