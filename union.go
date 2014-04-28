package recurrence

import (
	"encoding/json"
	"time"
)

// Computes the set union of a slice of Schedules.
type Union []Schedule

func (self Union) IsOccurring(t time.Time) bool {
	for _, r := range self {
		if r.IsOccurring(t) {
			return true
		}
	}
	return false
}

func (self Union) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(self)
}

func (self Union) MarshalJSON() ([]byte, error) {
	type faux Union
	return json.Marshal(struct {
		faux `json:"union"`
	}{faux: faux(self)})
}
