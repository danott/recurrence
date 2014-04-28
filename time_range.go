package recurrence

import (
	"encoding/json"
	"fmt"
	"time"
)

// A TimeRange represents a range of time, with a start and an end.
type TimeRange struct {
	Start time.Time
	End   time.Time
}

func (r TimeRange) IsOccurring(t time.Time) bool {
	return !(t.Before(r.Start) || t.After(r.End))
}

// Generate a TimeRange representing the entire year.
func YearRange(y int) TimeRange {
	return TimeRange{
		time.Date(y, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(y+1, time.January, 0, 0, 0, 0, 0, time.UTC),
	}
}

// Generate a TimeRange representing a specific month.
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

func (t TimeRange) Occurrences(other TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(other)
}

func (t TimeRange) occurrencesOfSchedule(s Schedule) chan time.Time {
	c := make(chan time.Time)

	go func() {
		for t := range t.eachDate() {
			if s.IsOccurring(t) {
				c <- t
			}
		}
		close(c)
	}()

	return c
}

func (r TimeRange) eachDate() chan time.Time {
	c := make(chan time.Time)

	go func() {
		for t := beginningOfDay(r.Start); !t.After(r.End); t = t.AddDate(0, 0, 1) {
			c <- t
		}
		close(c)
	}()

	return c
}

func (self *TimeRange) UnmarshalJSON(b []byte) error {
	var mixed interface{}
	var err error

	json.Unmarshal(b, &mixed)

	value, _ := mixed.(map[string]interface{})["start"]
	t, _ := time.Parse("2006-01-02", value.(string))
	self.Start = t

	value, _ = mixed.(map[string]interface{})["end"]
	t, _ = time.Parse("2006-01-02", value.(string))
	self.End = t

	return err
}

func beginningOfDay(t time.Time) time.Time {
	return t.Add(time.Hour * -12).Round(time.Hour * 24)
}
