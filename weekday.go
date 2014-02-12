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

func (w Weekday) Includes(t time.Time) bool {
	return t.Weekday() == time.Weekday(w)
}
