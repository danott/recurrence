package recurrence

import "time"

type Week int

func (w Week) IsOccurring(t time.Time) bool {
	if w := int(w); w == Last {
		return isLastWeekInMonth(t)
	} else {
		return weekInMonth(t) == w
	}
}

func (w Week) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(w)
}

func weekInMonth(t time.Time) int {
	return ((t.Day() - 1) / 7) + 1
}

func isLastWeekInMonth(t time.Time) bool {
	return t.Month() != t.AddDate(0, 0, 7).Month()
}
