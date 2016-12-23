package recurrence

import (
	"encoding/json"
	"fmt"
	"time"
)

// HourMinuteSecond encapsulates the hour, minute and second components of a time
type HourMinuteSecond struct {
	hour   int
	minute int
	second int
}

// NewHourMinuteSeconds is a convenience constructor method to create an instance of NewHourMinuteSeconds
func NewHourMinuteSeconds(h, m, s int) HourMinuteSecond {
	return HourMinuteSecond{h, m, s}
}

// Implement Schedule interface.
func (hms HourMinuteSecond) IsOccurring(t time.Time) bool {
	return int(hms.hour) == t.Hour() &&
		int(hms.minute) == t.Minute() &&
		int(hms.second) == t.Second()
}

func (hms HourMinuteSecond) Occurrences(tr TimeRange) []time.Time {
	return occurrencesFor(hms, tr)
}

func (hms HourMinuteSecond) nextAfter(t time.Time) (time.Time, error) {
	// Compare times ignoring nanoseconds
	thms := time.Date(t.Year(), t.Month(), t.Day(), hms.hour, hms.minute, hms.second, 0, time.UTC)
	t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, time.UTC)

	// This means this event is scheduled after the time that was passed
	if t.Before(thms) {
		return thms, nil
	}

	// Otherwise the event is scheduled for the next day
	return thms.AddDate(0, 0, 1), nil
}

// MarshalJSON returns a marshaled version of the underlying HourMinuteSecond instance
func (hms HourMinuteSecond) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"hour":   hms.hour,
		"minute": hms.minute,
		"second": hms.second,
	})
}

// UnmarshalJSON populates the *HourMinuteSecond pointer with values from the marshaled JSON bytes
func (hms *HourMinuteSecond) UnmarshalJSON(b []byte) error {
	m := map[string]interface{}{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	hourI, ok := m["hour"]
	if !ok {
		return fmt.Errorf("Missing 'hour' field")
	}
	if hour, ok := hourI.(float64); !ok {
		return fmt.Errorf("'hour' field should be integer")
	} else {
		hms.hour = int(hour)
	}

	minuteI, ok := m["minute"]
	if !ok {
		return fmt.Errorf("Missing 'minute' field")
	}
	if minute, ok := minuteI.(float64); !ok {
		return fmt.Errorf("'minute' field should be integer")
	} else {
		hms.minute = int(minute)
	}

	secondI, ok := m["second"]
	if !ok {
		return fmt.Errorf("Missing 'second' field")
	}
	if second, ok := secondI.(float64); !ok {
		return fmt.Errorf("'second' field should be integer")
	} else {
		hms.second = int(second)
	}

	return nil
}
