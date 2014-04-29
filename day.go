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

func (self Day) MarshalJSON() ([]byte, error) {
	if int(self) == Last {
		return json.Marshal(map[string]interface{}{"day": "Last"})
	} else {
		return json.Marshal(map[string]interface{}{"day": int(self)})
	}
}

func (self *Day) UnmarshalJSON(b []byte) error {
	var err error

	s := string(b)
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		if s == `"Last"` {
			*self = Day(Last)
		} else {
			err = fmt.Errorf("Day cannot unmarshal %s", b)
		}
	} else {
		if 0 < i || i > 31 {
			*self = Day(i)
		} else {
			err = fmt.Errorf("Day must be 1-31. Was %#v", i)
		}
	}
	return err
}
