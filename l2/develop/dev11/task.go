package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type CustomDate struct {
	time.Time
}

const layout = "2006-01-02"

func (c *CustomDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"") // remove quotes
	if s == "null" {
		return
	}
	c.Time, err = time.Parse(layout, s)
	return
}

func (c CustomDate) MarshalJSON() ([]byte, error) {
	if c.Time.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, c.Time.Format(layout))), nil
}

func parseEvent(body io.ReadCloser) Event {
	data, _ := ioutil.ReadAll(body)
	var event Event
	err := json.Unmarshal(data, &event)
	if err != nil {

	}
	return event
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	event := parseEvent(r.Body)
	w.Write(addEvent(&event))
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	event := parseEvent(r.Body)
	w.Write(updateEv(&event))
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	event := parseEvent(r.Body)
	w.Write(deleteEv(event.Name))
}

func eventsForDay(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	w.Write(getEventsForDay(date))
}

func eventsForWeek(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	w.Write(getEventsForWeek(date))
}

func eventsForMonth(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	w.Write(getEventsForMonth(date))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello there")
	})

	router.HandleFunc("/create_event", createEvent)
	router.HandleFunc("/update_event", updateEvent)
	router.HandleFunc("/delete_event", deleteEvent)
	router.HandleFunc("/events_for_day", eventsForDay)
	router.HandleFunc("/events_for_week", eventsForWeek)
	router.HandleFunc("/events_for_month", eventsForMonth)
	configuredRouter := LogMiddleware(router)

	err := http.ListenAndServe(":3333", configuredRouter)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
