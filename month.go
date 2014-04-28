package recurrence

import (
	"encoding/json"
	"time"
)

// A Month represents a month of the year. Just like time.Month.
type Month time.Month

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func (self Month) IsOccurring(t time.Time) bool {
	return t.Month() == time.Month(self)
}

func (self Month) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(self)
}

func (self Month) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"month": int(self)})
}
