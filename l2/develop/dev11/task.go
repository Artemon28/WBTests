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

func parseEvent(body io.ReadCloser) (Event, error) {
	data, _ := ioutil.ReadAll(body)
	var event Event
	err := json.Unmarshal(data, &event)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	event, err := parseEvent(r.Body)
	if err != nil {
		w.WriteHeader(400)
		_, err := w.Write([]byte("Error with parsing body"))
		if err != nil {
			w.WriteHeader(500)
		}
		return
	}
	result := addEvent(&event)
	if result[2] == 'e' {
		w.WriteHeader(503)
		return
	}
	w.WriteHeader(200)
	_, err = w.Write(result)
	if err != nil {
		w.WriteHeader(500)
	}
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	event, err := parseEvent(r.Body)
	if err != nil {
		w.WriteHeader(400)
		_, err := w.Write([]byte("Error with parsing body"))
		if err != nil {
			w.WriteHeader(500)
		}
		return
	}
	result := updateEv(&event)
	if result[2] == 'e' {
		w.WriteHeader(503)
		return
	}
	w.WriteHeader(200)
	_, err = w.Write(result)
	if err != nil {
		w.WriteHeader(500)
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	event, err := parseEvent(r.Body)
	if err != nil {
		w.WriteHeader(400)
		_, err := w.Write([]byte("Error with parsing body"))
		if err != nil {
			w.WriteHeader(500)
		}
		return
	}
	result := deleteEv(event.Name)
	if result[2] == 'e' {
		w.WriteHeader(503)
		return
	}
	w.WriteHeader(200)
	_, err = w.Write(result)
	if err != nil {
		w.WriteHeader(500)
	}
}

func eventsForDay(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	if date == "" {
		w.WriteHeader(400)
		_, err := w.Write([]byte("No field date in request"))
		if err != nil {
			w.WriteHeader(500)
		}
		return
	}
	result := getEventsForDay(date)
	if result[2] == 'e' {
		w.WriteHeader(503)
		return
	}
	w.WriteHeader(200)
	_, err := w.Write(result)
	if err != nil {
		w.WriteHeader(500)
	}
}

func eventsForWeek(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	if date == "" {
		w.WriteHeader(400)
		_, err := w.Write([]byte("No field date in request"))
		if err != nil {
			w.WriteHeader(500)
		}
		return
	}
	result := getEventsForWeek(date)
	if result[2] == 'e' {
		w.WriteHeader(503)
		return
	}
	w.WriteHeader(200)
	_, err := w.Write(result)
	if err != nil {
		w.WriteHeader(500)
	}
}

func eventsForMonth(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	if date == "" {
		w.WriteHeader(400)
		_, err := w.Write([]byte("No field date in request"))
		if err != nil {
			w.WriteHeader(500)
		}
		return
	}
	result := getEventsForMonth(date)
	if result[2] == 'e' {
		w.WriteHeader(503)
		return
	}
	w.WriteHeader(200)
	_, err := w.Write(result)
	if err != nil {
		w.WriteHeader(500)
	}
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
