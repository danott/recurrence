package recurrence

import (
	"encoding/json"
	"fmt"
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

func (self *Intersection) UnmarshalJSON(b []byte) error {
	var mixed interface{}

	json.Unmarshal(b, &mixed)

	switch mixed.(type) {
	case []interface{}:
		for _, value := range mixed.([]interface{}) {
			bytes, _ := json.Marshal(value)
			schedule, err := ScheduleUnmarshalJSON(bytes)
			if err != nil {
				return err
			}
			*self = append(*self, schedule)
		}
	default:
		return fmt.Errorf("intersection must be a slice")
	}

	return nil
}
