package recurrence

import (
	"encoding/json"
	"time"
)

// Computes the set union of a slice of Schedules.
type Union []Schedule

func (u Union) IsOccurring(t time.Time) bool {
	for _, r := range u {
		if r.IsOccurring(t) {
			return true
		}
	}
	return false
}

func (u Union) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(u)
}

func (u Union) MarshalJSON() ([]byte, error) {
	type faux Union
	return json.Marshal(struct {
		faux `json:"Union"`
	}{faux: faux(u)})
}
