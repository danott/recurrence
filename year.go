package recurrence

import (
	"encoding/json"
	"fmt"
	"time"
)

// Represents a year.
type Year int

// Implement Stringer interface.
func (self Year) String() string {
	return string(self)
}

// Implement Schedule interface.
func (self Year) IsOccurring(t time.Time) bool {
	return t.Year() == int(self)
}

// Implement Schedule interface.
func (self Year) Occurrences(tr TimeRange) []time.Time {
	return occurrencesFor(self, tr)
}

func (self Year) nextAfter(t time.Time) (time.Time, error) {
	desiredYear := int(self)

	if t.Year() == desiredYear && !isLastDayOfYear(t) {
		return t.AddDate(0, 0, 1), nil
	}

	if t.Year() < desiredYear {
		return time.Date(desiredYear, time.January, 1, 0, 0, 0, 0, time.UTC), nil
	}

	var zeroTime time.Time
	return zeroTime, fmt.Errorf("no more occurrences after %s", t)
}

// Implement json.Marshaler interface.
func (self Year) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"year": int(self)})
}

func isLastDayOfYear(t time.Time) bool {
	return t.Month() == time.December && t.Day() == 31
}
