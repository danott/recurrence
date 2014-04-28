package recurrence

import (
	"encoding/json"
	"time"
)

// A Day specifies a day of the month. (1, 2, 3, ...31)
type Day int

func (self Day) IsOccurring(t time.Time) bool {
	if self := int(self); self == Last {
		return isLastDayInMonth(t)
	} else {
		return self == t.Day()
	}
}

func (self Day) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(self)
}

func isLastDayInMonth(t time.Time) bool {
	return t.Month() != t.AddDate(0, 0, 1).Month()
}

func (self Day) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"day": int(self)})
}
