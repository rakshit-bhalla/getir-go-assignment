package utils

import (
	"testing"
	"time"
)

func TestStrToTime(t *testing.T) {
	date := "2022-01-09"
	got, err := StrToTime(date)
	if err != nil {
		t.Errorf("%q", err.Error())
	}
	want := time.Date(2022, 1, 9, 0, 0, 0, 0, time.UTC)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
