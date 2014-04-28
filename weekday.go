package recurrence

import (
	"encoding/json"
	"time"
)

// A Weekday represents a day of the week. (Sunday, Monday, ...Saturday)
type Weekday time.Weekday

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (self Weekday) IsOccurring(t time.Time) bool {
	return t.Weekday() == time.Weekday(self)
}

func (self Weekday) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(self)
}

func (self Weekday) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"weekday": int(self)})
}
