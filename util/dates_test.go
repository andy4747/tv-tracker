package util

import "testing"

func TestGetCurrentDate(t *testing.T) {
	date := GetCurrentDate()
	if date == "" {
		t.Errorf("couldn't get the current date err: %v\n", date)
	}
}
