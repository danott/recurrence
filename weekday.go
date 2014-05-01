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

func (self Weekday) Occurrences(tr TimeRange) chan time.Time {
	ch := make(chan time.Time)

	go func() {
		start := tr.Start.AddDate(0, 0, -1)
		end := tr.End
		for t, err := self.NextAfter(start); err == nil && !t.After(end); t, err = self.NextAfter(t) {
			ch <- t
		}
		close(ch)
	}()

	return ch
}

func (self Weekday) String() string {
	return time.Weekday(self).String()
}

func (self Weekday) NextAfter(t time.Time) (time.Time, error) {
	diff := int(self) - int(t.Weekday())
	if diff <= 0 {
		diff += 7
	}
	return t.AddDate(0, 0, diff), nil
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
