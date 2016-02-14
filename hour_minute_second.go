package recurrence

import (
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

func (hms HourMinuteSecond) Occurrences(tr TimeRange) chan time.Time {
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
