package recurrence

import "time"

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
