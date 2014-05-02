package recurrence

import "time"

// The Schedule interface is the foundation of the recurrence package. Types
// satisfying the Schedule interface can be used to determine if a time.Time
// occurs in the Schedule, and generate time.Times satisfying the Schedule.
type Schedule interface {
	IsOccurring(time.Time) bool
	Occurrences(TimeRange) chan time.Time
}

type nextable interface {
	nextAfter(time.Time) (time.Time, error)
}

func occurrencesFor(schedule nextable, timeRange TimeRange) chan time.Time {
	ch := make(chan time.Time)

	go func() {
		start := timeRange.Start.AddDate(0, 0, -1)
		end := timeRange.End
		for t, err := schedule.nextAfter(start); err == nil && !t.After(end); t, err = schedule.nextAfter(t) {
			if !t.After(end) {
				ch <- beginningOfDay(t)
			}
		}
		close(ch)
	}()

	return ch
}
