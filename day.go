package recurrence

import (
	"encoding/json"
	"time"
)

// A Day specifies a day of the month. (1, 2, 3, ...31)
type Day int

func (d Day) IsOccurring(t time.Time) bool {
	if d := int(d); d == Last {
		return isLastDayInMonth(t)
	} else {
		return d == t.Day()
	}
}

func (d Day) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(d)
}

func isLastDayInMonth(t time.Time) bool {
	return t.Month() != t.AddDate(0, 0, 1).Month()
}

func (d Day) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"Day": int(d)})
}
