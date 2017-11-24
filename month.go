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

// Implement Stringer interface.
func (m Month) String() string {
	return time.Month(m).String()
}

// Implement Schedule interface.
func (m Month) IsOccurring(t time.Time) bool {
	return t.Month() == time.Month(m)
}

// Implement Schedule interface.
func (m Month) Occurrences(tr TimeRange) chan time.Time {
	return occurrencesFor(m, tr)
}

func (m Month) nextAfter(t time.Time) (time.Time, error) {
	desiredMonth := int(m)
	tMonth := int(t.Month())

	if tMonth < desiredMonth {
		return time.Date(t.Year(), time.Month(desiredMonth), 1, 0, 0, 0, 0, time.UTC), nil
	}

	if tMonth > desiredMonth || (tMonth == desiredMonth && isLastDayInMonth(t)) {
		return time.Date(t.Year()+1, time.Month(desiredMonth), 1, 0, 0, 0, 0, time.UTC), nil
	}

	return t.AddDate(0, 0, 1), nil
}

// Implement json.Marshaler interface.
func (m Month) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"month": time.Month(m).String()})
}

// Implement json.Unmarshaler interface.
func (m *Month) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case `1`, `"January"`:
		*m = January
	case `2`, `"February"`:
		*m = February
	case `3`, `"March"`:
		*m = March
	case `4`, `"April"`:
		*m = April
	case `5`, `"May"`:
		*m = May
	case `6`, `"June"`:
		*m = June
	case `7`, `"July"`:
		*m = July
	case `8`, `"August"`:
		*m = August
	case `9`, `"September"`:
		*m = September
	case `10`, `"October"`:
		*m = October
	case `11`, `"November"`:
		*m = November
	case `12`, `"December"`:
		*m = December
	default:
		return fmt.Errorf("Weekday cannot unmarshal %s", b)
	}

	return nil
}
