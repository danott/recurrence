package recurrence

import "time"

// The Schedule interface is the foundation of the recurrence package. Types
// satisfying the Schedule interface can be used to determine if a time.Time
// occurs in the Schedule, and generate time.Times satisfying the Schedule.
type Schedule interface {
	IsOccurring(time.Time) bool
	Occurrences(TimeRange) chan time.Time
}
