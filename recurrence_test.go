package main

import (
	"testing"
	"time"
)

func TestDaily(t *testing.T) {
	d, _ := time.Parse("1/2/2006", "1/2/2006")
	expected, _ := time.Parse("1/2/2006", "1/3/2006")

	r := Recurrence{Daily}
	if r.NextDate(d) != expected {
		t.Errorf("Not as expected")
	}
}

func TestWeekly(t *testing.T) {
	d, _ := time.Parse("1/2/2006", "1/2/2006")
	expected, _ := time.Parse("1/2/2006", "1/9/2006")

	r := Recurrence{Weekly}
	if r.NextDate(d) != expected {
		t.Errorf("Not as expected")
	}
}

func TestMonthly(t *testing.T) {
	d, _ := time.Parse("1/2/2006", "1/2/2006")
	expected, _ := time.Parse("1/2/2006", "2/2/2006")

	r := Recurrence{Monthly}
	if r.NextDate(d) != expected {
		t.Errorf("Not as expected")
	}
}
