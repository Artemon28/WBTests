package main

import (
	"L0/cache"
	"L0/postgresDataBase"
	"context"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {
	//connection to DataBase
	connectionConfig := postgresDataBase.ConnectionConfig{Username: "postgres", Password: "UJBNVJ", Host: "localhost", Port: "5432", Database: "postgres"}
	connection, _ := postgresDataBase.NewClient(context.TODO(), connectionConfig)

	//Our cache map and recover of it
	orderCache, err := cache.Recover(connection)
	if err != nil {
		log.Println("Error with recovering cache from DB", err)
	}

	//connect to NATS streaming
	sc, err := stan.Connect("prod", "pub-test2")
	if err != nil {
		log.Println(err)
	}
	defer sc.Close()

	var mutex sync.Mutex
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		mutex.Lock()
		newOrder, err := cache.ConvertFromJson(m.Data)
		if err != nil {
			log.Println("error with trying to write data to the cache\n", err, "\n", string(m.Data))
			mutex.Unlock()
			return
		}
		orderCache[newOrder.Order_uid] = newOrder
		mutex.Unlock()
		err = cache.WriteToDB(connection, newOrder)
		if err != nil {
			log.Fatal("Error with writing data from canal to DB", err)
		}
	})
	if err != nil {
		log.Println(err)
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
				log.Println(err)
			}
		}(r.Body)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		_, ok := orderCache[string(body)]
		switch r.Method {
		case http.MethodPost:
			orderJson, err := json.Marshal(orderCache[string(body)])
			if err != nil {
				log.Println("Can't convert data from cache to JSON", err)
			}
			w.Header().Set("Content-Type", "application/json")
			if ok {
				w.WriteHeader(http.StatusOK)
				_, err2 := w.Write(orderJson)
				if err2 != nil {
					log.Println(err2)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				_, err3 := w.Write(nil)
				if err3 != nil {
					log.Println(err3)
				}
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Error with server at port 8080", err)
	}
	//nats-streaming-server -cid prod
}
