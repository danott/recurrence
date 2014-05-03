package recurrence

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// A Week represents a week of the month. This is most useful in combination
// with other entities satisfying the Schedule interface.
type Week int

// Implement Stringer interface.
func (self Week) String() string {
	switch int(self) {
	case 1:
		return "(First Week)"
	case 2:
		return "(Second Week)"
	case 3:
		return "(Third Week)"
	case 4:
		return "(Fourth Week)"
	case 5:
		return "(Fifth Week)"
	case Last:
		return "(Last Week)"
	default:
		return "(Never Week)"
	}
}

// Implement Schedule interface.
func (self Week) IsOccurring(t time.Time) bool {
	if weekInt := int(self); weekInt == Last {
		return isLastWeekInMonth(t)
	} else {
		return weekInMonth(t) == weekInt
	}
}

// Implement Schedule interface.
func (self Week) Occurrences(tr TimeRange) chan time.Time {
	return occurrencesFor(self, tr)
}

func (self Week) nextAfter(t time.Time) (time.Time, error) {
	desiredWeek := int(self)

	if desiredWeek == 1 {
		if t.Day() < 7 || isLastDayInMonth(t) {
			return t.AddDate(0, 0, 1), nil
		}

		return firstDayOfMonth(t).AddDate(0, 1, 0), nil
	}

	if desiredWeek == 2 {
		if t.Day() < 7 {
			return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, 7), nil
		}

		if t.Day() > 13 {
			return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, 7), nil
		}

		return t.AddDate(0, 0, 1), nil
	}

	if desiredWeek == 3 {
		if t.Day() < 14 {
			return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, 14), nil
		}

		if t.Day() > 20 {
			return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, 14), nil
		}

		return t.AddDate(0, 0, 1), nil
	}

	if desiredWeek == 4 {
		if t.Day() < 21 {
			return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, 21), nil
		}

		if t.Day() > 27 {
			return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, 21), nil
		}

		return t.AddDate(0, 0, 1), nil
	}

	if desiredWeek == 5 {
		totalDaysInMonth := lastDayOfMonth(t).Day()

		if totalDaysInMonth < 29 {
			return self.nextAfter(t.AddDate(0, 1, 0))
		}

		if isLastDayInMonth(t) {
			return self.nextAfter(t.AddDate(0, 0, 1))
		}

		if t.Day() < 27 {
			return firstDayOfMonth(t).AddDate(0, 0, 28), nil
		}

		return t.AddDate(0, 0, 1), nil
	}

	if desiredWeek == Last {
		totalDaysInMonth := lastDayOfMonth(t).Day()

		if isLastDayInMonth(t) {
			return self.nextAfter(t.AddDate(0, 0, 1))
		}

		if t.Day() < totalDaysInMonth-7 {
			return lastDayOfMonth(t).AddDate(0, 0, -6), nil
		}

		return t.AddDate(0, 0, 1), nil
	}

	return t, fmt.Errorf("You should never get here.")
}

// Implement json.Marshaler interface.
func (self Week) MarshalJSON() ([]byte, error) {
	if int(self) == Last {
		return json.Marshal(map[string]interface{}{"week": "Last"})
	} else {
		return json.Marshal(map[string]interface{}{"week": int(self)})
	}
}

// Implement json.Unmarshaler interface.
func (self *Week) UnmarshalJSON(b []byte) error {
	switch s := string(b); s {
	case `1`, `2`, `3`, `4`, `5`:
		i, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			return err
		}
		*self = Week(i)
	case `"Last"`:
		*self = Week(Last)
	default:
		return fmt.Errorf("Week cannot unmarshal %s", b)
	}

	return nil
}

func isLastWeekInMonth(t time.Time) bool {
	return t.Month() != t.AddDate(0, 0, 7).Month()
}

func weekInMonth(t time.Time) int {
	return ((t.Day() - 1) / 7) + 1
}
