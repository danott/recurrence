package recurrence

import (
	"encoding/json"
	"fmt"
	"time"
)

// A Week represents a week of the month. This is most useful in combination
// with other entities satisfying the Schedule interface.
type Week int

func (self Week) IsOccurring(t time.Time) bool {
	if weekInt := int(self); weekInt == Last {
		return isLastWeekInMonth(t)
	} else {
		return weekInMonth(t) == weekInt
	}
}

func (self Week) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(self)
}

func weekInMonth(t time.Time) int {
	return ((t.Day() - 1) / 7) + 1
}

func isLastWeekInMonth(t time.Time) bool {
	return t.Month() != t.AddDate(0, 0, 7).Month()
}

func (self Week) MarshalJSON() ([]byte, error) {
	if int(self) == Last {
		return json.Marshal(map[string]interface{}{"week": "Last"})
	} else {
		return json.Marshal(map[string]interface{}{"week": int(self)})
	}
}

func (self *Week) UnmarshalJSON(b []byte) error {
	var err error

	switch string(b) {
	case `1`:
		*self = Week(1)
	case `2`:
		*self = Week(2)
	case `3`:
		*self = Week(3)
	case `4`:
		*self = Week(4)
	case `5`:
		*self = Week(5)
	case `"Last"`:
		*self = Week(Last)
	default:
		err = fmt.Errorf("Week cannot unmarshal %s", b)
	}
	return err
}
