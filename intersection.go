package recurrence

import (
	"encoding/json"
	"time"
)

// Computes the set intersection of a slice of Schedules.
type Intersection []Schedule

func (i Intersection) IsOccurring(t time.Time) bool {
	for _, r := range i {
		if r.IsOccurring(t) == false {
			return false
		}
	}
	return true
}

func (i Intersection) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(i)
}

func (i Intersection) MarshalJSON() ([]byte, error) {
	type faux Intersection
	return json.Marshal(struct {
		faux `json:"Intersection"`
	}{faux: faux(i)})
}
