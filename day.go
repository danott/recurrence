package recurrence

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// A Day specifies a day of the month. (1, 2, 3, ...31)
type Day int

func (self Day) IsOccurring(t time.Time) bool {
	if self := int(self); self == Last {
		return isLastDayInMonth(t)
	} else {
		return self == t.Day()
	}
}

func (self Day) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(self)
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

func (self Day) MarshalJSON() ([]byte, error) {
	if int(self) == Last {
		return json.Marshal(map[string]interface{}{"day": "Last"})
	} else {
		return json.Marshal(map[string]interface{}{"day": int(self)})
	}
}

func (self *Day) UnmarshalJSON(b []byte) error {
	s := string(b)

	i, err := strconv.ParseInt(s, 10, 0)

	if err != nil {
		if s != `"Last"` {
			return fmt.Errorf("day cannot unmarshal %s", b)
		} else {
			*self = Day(Last)
		}
	} else {
		if i < 1 || i > 31 {
			return fmt.Errorf("day must be 1-31. Was %#v", i)
		} else {
			*self = Day(i)
		}
	}

	return nil
}
