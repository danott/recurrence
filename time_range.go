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

// Implement Schedule interface.
func (self TimeRange) IsOccurring(t time.Time) bool {
	return !(t.Before(self.Start) || t.After(self.End))
}

// Implement Schedule interface.
func (self TimeRange) Occurrences(other TimeRange) chan time.Time {
	return occurrencesFor(self, other)
}

func (self TimeRange) nextAfter(t time.Time) (time.Time, error) {
	if t.Before(self.Start) {
		return self.Start, nil
	}

	if t.Before(self.End) {
		return t.AddDate(0, 0, 1), nil
	}

	var zeroTime time.Time
	return zeroTime, fmt.Errorf("no more occurrences after %s", t)
}

// Implement json.Unmarshaler interface.
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

// NewTimeRange let's you create a new TimeRange from the time format "2006-01-02"
func NewTimeRange(start, end string) TimeRange {
	tStart, err := time.Parse("2006-01-02", start)
	if err != nil {
		panic(`NewDate(string) requires format "2006-01-02"`)
	}

	tEnd, err := time.Parse("2006-01-02", end)
	if err != nil {
		panic(`NewDate(string) requires format "2006-01-02"`)
	}

	return TimeRange{tStart, tEnd}
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

func beginningOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

func (self TimeRange) eachDate() chan time.Time {
	c := make(chan time.Time)

	go func() {
		for t := beginningOfDay(self.Start); !t.After(self.End); t = t.AddDate(0, 0, 1) {
			c <- t
		}
		close(c)
	}()

	return c
}
