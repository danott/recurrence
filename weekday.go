package recurrence

import "time"

const (
	Sunday time.Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Weekday struct {
	weekday time.Weekday
}

func (w Weekday) Includes(t time.Time) bool {
	return t.Weekday() == w.weekday
}
