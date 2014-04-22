package recurrence

import (
	"encoding/json"
	"time"
)

// A Week represents a week of the month. This is most useful in combination
// with other entities satisfying the Schedule interface.
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

func (w Week) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"week": int(w)})
}
