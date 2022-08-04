package main

import (
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	expectedTime := time.Now()
	resultTime := getTime()

	if resultTime.Equal(expectedTime) {
		t.Errorf("Time now and time from ntp lib not the same")
	}
}
