package main

import (
	"fmt"
	"time"
)

const (
	Daily   = iota
	Weekly  = iota
	Monthly = iota
	Yearly  = iota
)

type Recurrence struct {
	Frequency int8
}

func (r Recurrence) NextDate(d time.Time) (nextDate time.Time) {
	switch r.Frequency {
	case Daily:
		nextDate = d.AddDate(0, 0, 1)
	case Weekly:
		nextDate = d.AddDate(0, 0, 7)
	case Monthly:
		nextDate = d.AddDate(0, 1, 0)
	case Yearly:
		nextDate = d.AddDate(1, 0, 0)
	}

	return
}

func (r Recurrence) GenerateDays(d time.Time, end time.Time) chan time.Time {
	c := make(chan time.Time)

	go func() {
		for d.Before(end) {
			d = r.NextDate(d)
			c <- d
		}
		close(c)
	}()

	return c
}

func main() {
	rd := Recurrence{Daily}
	start := time.Now().Truncate(time.Hour)
	end := time.Now().Truncate(time.Hour).AddDate(2, 0, 0)

	for d := range rd.GenerateDays(start, end) {
		fmt.Println(d)
	}
}
