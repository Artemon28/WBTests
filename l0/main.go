package main

import (
	"L0/cache"
	"L0/postgresDataBase"
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//connection to DataBase
	connectionConfig := postgresDataBase.ConnectionConfig{Username: "postgres", Password: "UJBNVJ", Host: "localhost", Port: "5432", Database: "postgres"}
	connection, _ := postgresDataBase.NewClient(context.TODO(), connectionConfig)

	//Our cache map and recover of it
	orderCache, err := cache.Recover(connection)
	if err != nil {
		fmt.Println("Error with recovering cache from DB", err)
	}

	//connect to NATS streaming
	sc, err := stan.Connect("prod", "pub-test2")
	if err != nil {
		fmt.Println(err)
	}
	defer sc.Close()

	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		newOrder, err := cache.ConvertFromJson(m.Data)
		if err != nil {
			fmt.Print("error with trying to write data to the cache\n", err, "\n", string(m.Data))
		}
		orderCache[newOrder.Order_uid] = newOrder
		err = cache.WriteToDB(connection, newOrder)
		if err != nil {
			log.Fatal("Error with writing data from canal to DB", err)
		}
	})
	if err != nil {
		fmt.Println(err)
	}
	defer sub.Close()

	//Http server connection
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "L0/index.html")
	})

	http.HandleFunc("/order/", func(w http.ResponseWriter, r *http.Request) {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(r.Body)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		_, ok := orderCache[string(body)]
		switch r.Method {
		case http.MethodPost:
			orderJson, err := json.Marshal(orderCache[string(body)])
			if err != nil {
				fmt.Println("Can't convert data from cache to JSON", err)
			}
			w.Header().Set("Content-Type", "application/json")
			if ok {
				w.WriteHeader(http.StatusOK)
				_, err2 := w.Write(orderJson)
				if err2 != nil {
					fmt.Println(err2)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				_, err3 := w.Write(nil)
				if err3 != nil {
					fmt.Println(err3)
				}
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error with server at port 8080", err)
	}
	//nats-streaming-server -cid prod
}
