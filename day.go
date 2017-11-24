package recurrence

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// A Day specifies a day of the month. (1, 2, 3, ...31)
type Day int

// Implement Schedule interface.
func (d Day) IsOccurring(t time.Time) bool {
	if d := int(d); d == Last {
		return isLastDayInMonth(t)
	} else {
		return d == t.Day()
	}
}

// Implement Schedule interface.
func (d Day) Occurrences(tr TimeRange) chan time.Time {
	return occurrencesFor(d, tr)
}

func (d Day) nextAfter(t time.Time) (time.Time, error) {
	desiredDay := int(d)

	if desiredDay == Last {
		if isLastDayInMonth(t) {
			return t.AddDate(0, 0, 1).AddDate(0, 1, -1), nil
		}

		return firstDayOfMonth(t).AddDate(0, 2, -1), nil
	}

	if t.Day() > desiredDay {
		if isLastDayInMonth(t) && desiredDay == First {
			return t.AddDate(0, 0, 1), nil
		}

		return d.nextAfter(t.AddDate(0, 0, 1))
	}

	if t.Day() < desiredDay {
		totalDays := lastDayOfMonth(t).Day()
		if totalDays < desiredDay {
			return d.nextAfter(t.AddDate(0, 1, 0))
		}

		return time.Date(t.Year(), t.Month(), desiredDay, 0, 0, 0, 0, time.UTC), nil
	}

	totalDaysNextMonth := lastDayOfMonth(lastDayOfMonth(t).AddDate(0, 0, 1)).Day()
	if totalDaysNextMonth < desiredDay {
		return d.nextAfter(t.AddDate(0, 2, -1))
	}

	return t.AddDate(0, 1, 0), nil
}

// Implement json.Marshaler interface.
func (d Day) MarshalJSON() ([]byte, error) {
	if int(d) == Last {
		return json.Marshal(map[string]interface{}{"day": "Last"})
	} else {
		return json.Marshal(map[string]interface{}{"day": int(d)})
	}
}

// Implement json.Unmarshaler interface.
func (d *Day) UnmarshalJSON(b []byte) error {
	s := string(b)

	i, err := strconv.ParseInt(s, 10, 0)

	if err != nil {
		if s != `"Last"` {
			return fmt.Errorf("day cannot unmarshal %s", b)
		} else {
			*d = Day(Last)
		}
	} else {
		if i < 1 || i > 31 {
			return fmt.Errorf("day must be 1-31. Was %#v", i)
		} else {
			*d = Day(i)
		}
	}

	return nil
}

func isLastDayInMonth(t time.Time) bool {
	return t.Month() != t.AddDate(0, 0, 1).Month()
}

func lastDayOfMonth(t time.Time) time.Time {
	return firstDayOfMonth(t).AddDate(0, 1, -1)
}

func firstDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
}
