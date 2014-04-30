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
