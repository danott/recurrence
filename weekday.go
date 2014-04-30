package recurrence

import (
	"encoding/json"
	"fmt"
	"time"
)

// A Weekday represents a day of the week. (Sunday, Monday, ...Saturday)
type Weekday time.Weekday

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (self Weekday) IsOccurring(t time.Time) bool {
	return t.Weekday() == time.Weekday(self)
}

func (self Weekday) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(self)
}

func (self *Weekday) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case `0`, `"Sunday"`:
		*self = Sunday
	case `1`, `"Monday"`:
		*self = Monday
	case `2`, `"Tuesday"`:
		*self = Tuesday
	case `3`, `"Wednesday"`:
		*self = Wednesday
	case `4`, `"Thursday"`:
		*self = Thursday
	case `5`, `"Friday"`:
		*self = Friday
	case `6`, `"Saturday"`:
		*self = Saturday
	default:
		return fmt.Errorf("Weekday cannot unmarshal %s", b)
	}

	return nil
}

func (self Weekday) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"weekday": time.Weekday(self).String()})
}
