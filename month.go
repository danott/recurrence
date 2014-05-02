package recurrence

import (
	"encoding/json"
	"fmt"
	"time"
)

// A Month represents a month of the year. Just like time.Month.
type Month time.Month

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func (self Month) IsOccurring(t time.Time) bool {
	return t.Month() == time.Month(self)
}

func (self Month) Occurrences(tr TimeRange) chan time.Time {
	return occurrencesFor(self, tr)
}

func (self Month) nextAfter(t time.Time) (time.Time, error) {
	desiredMonth := int(self)
	tMonth := int(t.Month())

	if tMonth < desiredMonth {
		return time.Date(t.Year(), time.Month(desiredMonth), 1, 0, 0, 0, 0, time.UTC), nil
	}

	if tMonth > desiredMonth || (tMonth == desiredMonth && isLastDayInMonth(t)) {
		return time.Date(t.Year()+1, time.Month(desiredMonth), 1, 0, 0, 0, 0, time.UTC), nil
	}

	return t.AddDate(0, 0, 1), nil
}

func (self Month) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"month": time.Month(self).String()})
}

func (self *Month) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case `1`, `"January"`:
		*self = January
	case `2`, `"February"`:
		*self = February
	case `3`, `"March"`:
		*self = March
	case `4`, `"April"`:
		*self = April
	case `5`, `"May"`:
		*self = May
	case `6`, `"June"`:
		*self = June
	case `7`, `"July"`:
		*self = July
	case `8`, `"August"`:
		*self = August
	case `9`, `"September"`:
		*self = September
	case `10`, `"October"`:
		*self = October
	case `11`, `"November"`:
		*self = November
	case `12`, `"December"`:
		*self = December
	default:
		return fmt.Errorf("Weekday cannot unmarshal %s", b)
	}

	return nil
}

func (self Month) String() string {
	return time.Month(self).String()
}
