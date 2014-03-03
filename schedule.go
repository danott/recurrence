package recurrence

import "time"

type Schedule interface {
	IsOccurring(time.Time) bool
	Occurrences(TimeRange) chan time.Time
}
