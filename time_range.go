package recurrence

import (
	"fmt"
	"time"
)

type TimeRange struct {
	start time.Time
	end   time.Time
}

func (r TimeRange) Includes(t time.Time) bool {
	return !(t.Before(r.start) || t.After(r.end))
}

func YearRange(y int) TimeRange {
	return TimeRange{
		time.Date(y, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(y+1, time.January, 0, 0, 0, 0, 0, time.UTC),
	}
}

func MonthRange(month interface{}, year int) TimeRange {
	var m time.Month

	switch t := month.(type) {
	case int:
		m = time.Month(t)
	case Month:
		m = time.Month(t)
	case time.Month:
		m = t
	default:
		panic(fmt.Sprintf("MonthRange can't use %T", month))
	}

	return TimeRange{
		time.Date(year, m, 1, 0, 0, 0, 0, time.UTC),
		time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC),
	}
}

func Dates(t TimeRange, r Rule) (result []time.Time) {
	for t := range t.eachDate() {
		if r.Includes(t) {
			result = append(result, t)
		}
	}
	return
}

func (r TimeRange) eachDate() chan time.Time {
	c := make(chan time.Time)

	go func() {
		for t := r.start; !t.After(r.end); t = t.AddDate(0, 0, 1) {
			c <- t
		}
		close(c)
	}()

	return c
}
