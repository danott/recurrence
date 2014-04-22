package recurrence

import (
	"encoding/json"
	"time"
)

// Computes the set difference of two Schedules.
type Exclusion struct {
	Schedule Schedule `json:"schedule"`
	Exclude  Schedule `json:"exclude"`
}

func (d Exclusion) IsOccurring(t time.Time) bool {
	if d.Exclude.IsOccurring(t) {
		return false
	}

	if d.Schedule.IsOccurring(t) {
		return true
	}

	return false
}

func (d Exclusion) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(d)
}

func (d Exclusion) MarshalJSON() ([]byte, error) {
	type faux Exclusion
	return json.Marshal(struct {
		faux `json:"exclusion"`
	}{faux: faux(d)})
}
