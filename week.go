package recurrence

import (
	"encoding/json"
	"time"
)

// A Week represents a week of the month. This is most useful in combination
// with other entities satisfying the Schedule interface.
type Week int

func (self Week) IsOccurring(t time.Time) bool {
	if weekInt := int(self); weekInt == Last {
		return isLastWeekInMonth(t)
	} else {
		return weekInMonth(t) == weekInt
	}
}

func (self Week) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(self)
}

func weekInMonth(t time.Time) int {
	return ((t.Day() - 1) / 7) + 1
}

func isLastWeekInMonth(t time.Time) bool {
	return t.Month() != t.AddDate(0, 0, 7).Month()
}

func (self Week) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"week": int(self)})
}
