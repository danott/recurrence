package recurrence

import "time"

type Rule interface {
	Includes(time.Time) bool
	Dates(TimeRange) chan time.Time
}
