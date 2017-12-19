package recurrence

import (
	"encoding/json"
	"fmt"
	"time"
)

// Year represents a year.
type Year int

// implements the Stringer interface.
func (y Year) String() string {
	return string(y)
}

// IsOccurring implements the Schedule interface.
func (y Year) IsOccurring(t time.Time) bool {
	return t.Year() == int(y)
}

// Occurrences implements the Schedule interface.
func (y Year) Occurrences(tr TimeRange) chan time.Time {
	return occurrencesFor(y, tr)
}

func (y Year) nextAfter(t time.Time) (time.Time, error) {
	desiredYear := int(y)

	if t.Year() == desiredYear && !isLastDayOfYear(t) {
		return t.AddDate(0, 0, 1), nil
	}

	if t.Year() < desiredYear {
		return time.Date(desiredYear, time.January, 1, 0, 0, 0, 0, time.UTC), nil
	}

	var zeroTime time.Time
	return zeroTime, fmt.Errorf("no more occurrences after %s", t)
}

// MarshalJSON implements the json.Marshaler interface.
func (y Year) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"year": int(y)})
}

func isLastDayOfYear(t time.Time) bool {
	return t.Month() == time.December && t.Day() == 31
}
