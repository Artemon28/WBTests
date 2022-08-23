package main

import (
	"strings"
	"time"
)

type Event struct {
	Name string `json:"name"`
	Date CustomDate
}

var calendar = make(map[string]*Event, 1000)

func addEvent(event *Event) []byte {
	calendar[event.Name] = event
	return []byte(`{"result": "success added"}`)
}

func updateEv(event *Event) []byte {
	if _, ok := calendar[event.Name]; ok {
		calendar[event.Name] = event
		return []byte(`{"result": "success updated"}`)
	}
	return []byte(`{"error": "don't have this event to update it'"}`)
}

func deleteEv(name string) []byte {
	delete(calendar, name)
	return []byte(`{"result": "success deleted"}`)
}

func getEventsForDay(dateStr string) []byte {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return []byte(`{"error": "` + err.Error() + `"}`)
	}
	eventsThisDay := strings.Builder{}
	for _, v := range calendar {
		if v.Date.Time == date {
			eventsThisDay.WriteString(v.Name)
			eventsThisDay.WriteString(" ")
		}
	}
	return []byte(`{"result": "` + eventsThisDay.String() + `"}`)
}

func getEventsForWeek(dateStr string) []byte {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return []byte(`{"error": "` + err.Error() + `"}`)
	}
	eventsThisDay := strings.Builder{}
	for _, v := range calendar {
		vWeek, vYear := v.Date.ISOWeek()
		week, year := date.ISOWeek()
		if vWeek == week && vYear == year {
			eventsThisDay.WriteString(v.Name)
			eventsThisDay.WriteString(" ")
		}
	}
	return []byte(`{"result": "` + eventsThisDay.String() + `"}`)
}

func getEventsForMonth(dateStr string) []byte {
	date, err := time.Parse("2006-01", dateStr)
	if err != nil {
		return []byte(`{"error": "` + err.Error() + `"}`)
	}
	eventsThisDay := strings.Builder{}
	for _, v := range calendar {
		if v.Date.Month() == date.Month() && v.Date.Year() == date.Year() {
			eventsThisDay.WriteString(v.Name)
			eventsThisDay.WriteString(" ")
		}
	}
	return []byte(`{"result": "` + eventsThisDay.String() + `"}`)
}
