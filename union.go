package recurrence

import (
	"encoding/json"
	"fmt"
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

func (self *Union) UnmarshalJSON(b []byte) error {
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
		return fmt.Errorf("union must be a slice")
	}

	return nil
}
