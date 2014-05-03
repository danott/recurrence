package recurrence

import (
	"encoding/json"
	"fmt"
	"time"
)

// Computes the set intersection of a slice of Schedules.
type Intersection []Schedule

// Implement Schedule interface.
func (self Intersection) IsOccurring(t time.Time) bool {
	for _, r := range self {
		if r.IsOccurring(t) == false {
			return false
		}
	}

	return true
}

// Implement Schedule interface.
func (self Intersection) Occurrences(t TimeRange) chan time.Time {
	ch := make(chan time.Time)
	done := make(chan bool, len(self))
	candidates := make(chan time.Time)

	for _, schedule := range self {
		go func(schedule Schedule) {
			for t := range schedule.Occurrences(t) {
				candidates <- t
			}
			done <- true
		}(schedule)
	}

	go func() {
		candidatesMap := make(map[string]int)
		for candidate := range candidates {
			key := candidate.Format("20060102")
			foundCount, _ := candidatesMap[key]
			newFoundCount := foundCount + 1
			candidatesMap[key] = newFoundCount
			if newFoundCount == len(self) {
				ch <- candidate
			}
		}
	}()

	go func() {
		for i := 0; i < len(self); i++ {
			<-done
		}
		close(ch)
		close(done)
		close(candidates)
	}()

	return ch
}

// Implement json.Marshaler interface.
func (self Intersection) MarshalJSON() ([]byte, error) {
	type faux Intersection
	return json.Marshal(struct {
		faux `json:"intersection"`
	}{faux: faux(self)})
}

// Implement json.Unmarshaler interface.
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
