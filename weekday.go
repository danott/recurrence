package recurrence

import (
	"encoding/json"
	"fmt"
	"time"
)

// A Weekday represents a day of the week. (Sunday, Monday, ...Saturday)
type Weekday time.Weekday

// The days of the week
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// implements the Stringer interface.
func (w Weekday) String() string {
	return time.Weekday(w).String()
}

// IsOccurring implements the Schedule interface.
func (w Weekday) IsOccurring(t time.Time) bool {
	return t.Weekday() == time.Weekday(w)
}

// Occurrences implements the Schedule interface.
func (w Weekday) Occurrences(tr TimeRange) chan time.Time {
	return occurrencesFor(w, tr)
}

func (w Weekday) nextAfter(t time.Time) (time.Time, error) {
	diff := int(w) - int(t.Weekday())
	if diff <= 0 {
		diff += 7
	}
	return t.AddDate(0, 0, diff), nil
}

// MarshalJSON implements the json.Marshaler interface.
func (w Weekday) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"weekday": time.Weekday(w).String()})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (w *Weekday) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case `0`, `"Sunday"`:
		*w = Sunday
	case `1`, `"Monday"`:
		*w = Monday
	case `2`, `"Tuesday"`:
		*w = Tuesday
	case `3`, `"Wednesday"`:
		*w = Wednesday
	case `4`, `"Thursday"`:
		*w = Thursday
	case `5`, `"Friday"`:
		*w = Friday
	case `6`, `"Saturday"`:
		*w = Saturday
	default:
		return fmt.Errorf("Weekday cannot unmarshal %s", b)
	}

	return nil
}
