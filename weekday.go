package recurrence

import "time"

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

func (w Weekday) IsOccurring(t time.Time) bool {
	return t.Weekday() == time.Weekday(w)
}

func (w Weekday) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(w)
}
