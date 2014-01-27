package recurrence

import "time"

type Day struct {
	day int
}

func (d Day) Includes(t time.Time) bool {
	if d.day == Last {
		return isLastDayInMonth(t)
	} else {
		return t.Day() == d.day
	}
}

func isLastDayInMonth(t time.Time) bool {
	return t.Month() != t.AddDate(0, 0, 1).Month()
}
