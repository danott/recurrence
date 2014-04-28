package recurrence

import (
	"encoding/json"
	"time"
)

// Computes the set intersection of a slice of Schedules.
type Intersection []Schedule

func (self Intersection) IsOccurring(t time.Time) bool {
	for _, r := range self {
		if r.IsOccurring(t) == false {
			return false
		}
	}
	return true
}

func (self Intersection) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(self)
}

func (self Intersection) MarshalJSON() ([]byte, error) {
	type faux Intersection
	return json.Marshal(struct {
		faux `json:"intersection"`
	}{faux: faux(self)})
}
