package recurrence

import "time"

type OrdinalWeekday struct {
	week    int
	weekday time.Weekday
}

func (d OrdinalWeekday) Includes(t time.Time) bool {
	if d.week == Last {
		return d.weekday == t.Weekday() && isLastWeekInMonth(t)
	} else {
		return d.weekday == t.Weekday() && weekFromMonthStart(t) == d.week
	}
}

func weekFromMonthStart(t time.Time) int {
	return ((t.Day() - 1) / 7) + 1
}

func isLastWeekInMonth(t time.Time) bool {
	return t.Month() != t.AddDate(0, 0, 7).Month()
}
