package recurrence

import (
	"encoding/json"
	"fmt"
	"time"
)

// Represents a year.
type Year int

func (self Year) IsOccurring(t time.Time) bool {
	return t.Year() == int(self)
}

func (self Year) Occurrences(tr TimeRange) chan time.Time {
	ch := make(chan time.Time)
	go func() {
		for na, ok := self.NextAfter(tr.Start.AddDate(0, 0, -1)); ok == nil && !na.After(tr.End); na, ok = self.NextAfter(na) {
			ch <- na
		}
		close(ch)
	}()
	return ch
}

func (self Year) NextAfter(t time.Time) (time.Time, error) {
	if t.Year() > int(self) || (t.Year() == int(self) && t.Month() == time.December && t.Day() == 31) {
		var zeroTime time.Time
		return zeroTime, fmt.Errorf("never happens again")
	} else if t.Year() < int(self) {
		return time.Date(int(self), time.January, 1, 0, 0, 0, 0, time.UTC), nil
	} else {
		return t.AddDate(0, 0, 1), nil
	}
}

func (self Year) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"year": int(self)})
}
