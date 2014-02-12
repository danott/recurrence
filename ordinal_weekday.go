package recurrence

import "time"

type OrdinalWeekday struct {
	week    int
	weekday Weekday
}

func (o OrdinalWeekday) Includes(t time.Time) bool {
	return o.weekday.Includes(t) && weekMatches(o, t)
}

func weekMatches(o OrdinalWeekday, t time.Time) bool {
	if o.week == Last {
		return isLastWeekInMonth(t)
	} else {
		return weekInMonth(t) == o.week
	}
}

func weekInMonth(t time.Time) int {
	return ((t.Day() - 1) / 7) + 1
}

func isLastWeekInMonth(t time.Time) bool {
	return t.Month() != t.AddDate(0, 0, 7).Month()
}
