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

func (self Exclusion) IsOccurring(t time.Time) bool {
	if self.Exclude.IsOccurring(t) {
		return false
	}

	if self.Schedule.IsOccurring(t) {
		return true
	}

	return false
}

func (self Exclusion) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(self)
}

func (self Exclusion) MarshalJSON() ([]byte, error) {
	type faux Exclusion
	return json.Marshal(struct {
		faux `json:"exclusion"`
	}{faux: faux(self)})
}
