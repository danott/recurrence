package recurrence

import "time"

// The Schedule interface is the foundation of the recurrence package. Types
// satisfying the Schedule interface can be used to determine if a time.Time
// occurs in the Schedule, and generate time.Times satisfying the Schedule.
type Schedule interface {
	IsOccurring(time.Time) bool
	Occurrences(TimeRange) []time.Time
}

// @todo why is this not in the schedule interface?
type nextable interface {
	nextAfter(time.Time) (time.Time, error)
}

func occurrencesFor(schedule nextable, timeRange TimeRange) []time.Time {
	ts := make([]time.Time, 0)
	start := timeRange.Start.AddDate(0, 0, -1)
	end := timeRange.End

	for t, err := schedule.nextAfter(start); err == nil && !t.After(end); t, err = schedule.nextAfter(t) {
		if !t.After(end) {
			ts = append(ts, beginningOfDay(t))
		}
	}

	return ts
}
