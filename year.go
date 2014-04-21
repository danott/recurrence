package recurrence

import "time"

// Represents a year.
type Year int

func (y Year) IsOccurring(t time.Time) bool {
	return t.Year() == int(y)
}

func (y Year) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(y)
}
