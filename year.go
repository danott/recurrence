package recurrence

import (
	"encoding/json"
	"time"
)

// Represents a year.
type Year int

func (self Year) IsOccurring(t time.Time) bool {
	return t.Year() == int(self)
}

func (self Year) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(self)
}

func (self Year) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"year": int(self)})
}
