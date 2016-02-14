package recurrence

import (
	"encoding/json"
	"fmt"
	"time"
)

// Computes the set union of a slice of Schedules.
type Union []Schedule

// Implement Schedule interface.
func (self Union) IsOccurring(t time.Time) bool {
	for _, r := range self {
		if r.IsOccurring(t) {
			return true
		}
	}

	return false
}

// Implement Schedule interface.
func (self Union) Occurrences(t TimeRange) []time.Time {
	done := make(chan bool, len(self))
	candidates := make(chan time.Time, 100)
	ts := make([]time.Time, 0)

	for _, schedule := range self {
		go func(schedule Schedule) {
			for _, oc := range schedule.Occurrences(t) {
				candidates <- oc
			}
			done <- true
		}(schedule)
	}

	candidatesMap := make(map[string]bool)
	parallelDone := 0
	for parallelDone < len(self) {
		select {
		case selected := <-candidates:
			key := selected.Format("20060102")
			_, found := candidatesMap[key]
			if !found {
				candidatesMap[key] = true
				ts = append(ts, selected)
			}
		case <-done:
			parallelDone++
		}
	}

	// We can safely close the "done" channel
	close(done)

	// We must make sure we have fully drained the "candidates" channel
	stillLoop := true
	for stillLoop {
		select {
		case selected := <-candidates:
			key := selected.Format("20060102")
			_, found := candidatesMap[key]
			if !found {
				candidatesMap[key] = true
				ts = append(ts, selected)
			}
		default:
			stillLoop = false
		}
	}

	// We can also safely close the "candidates" channel
	close(candidates)

	return ts
}

// Implement json.Marshaler interface.
func (self Union) MarshalJSON() ([]byte, error) {
	type faux Union
	return json.Marshal(struct {
		faux `json:"union"`
	}{faux: faux(self)})
}

// Implement json.Unmarshaler interface.
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
