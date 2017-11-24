package recurrence

import (
	"encoding/json"
	"fmt"
	"time"
)

// Computes the set union of a slice of Schedules.
type Union []Schedule

// Implement Schedule interface.
func (u Union) IsOccurring(t time.Time) bool {
	for _, r := range u {
		if r.IsOccurring(t) {
			return true
		}
	}

	return false
}

// Implement Schedule interface.
func (u Union) Occurrences(t TimeRange) chan time.Time {
	ch := make(chan time.Time)
	done := make(chan bool, len(u))
	candidates := make(chan time.Time)

	for _, schedule := range u {
		go func(schedule Schedule) {
			for t := range schedule.Occurrences(t) {

				candidates <- t
			}
			done <- true
		}(schedule)
	}

	go func() {
		candidatesMap := make(map[string]bool)
		for candidate := range candidates {
			key := candidate.Format("20060102")
			_, found := candidatesMap[key]
			if !found {
				candidatesMap[key] = true
				ch <- candidate
			}
		}
	}()

	go func() {
		for i := 0; i < len(u); i++ {
			<-done
		}
		close(ch)
		close(done)
	}()

	return ch
}

// Implement json.Marshaler interface.
func (u Union) MarshalJSON() ([]byte, error) {
	type wrapper struct {
		Union []Schedule `json:"union"`
	}
	return json.Marshal(wrapper{Union: u})
}

// Implement json.Unmarshaler interface.
func (u *Union) UnmarshalJSON(b []byte) error {
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
			*u = append(*u, schedule)
		}
	default:
		return fmt.Errorf("union must be a slice")
	}

	return nil
}
