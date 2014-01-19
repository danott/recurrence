package recurrence

import "time"

type TimeRange struct {
	startTime time.Time
	endTime   time.Time
}

func (r TimeRange) Includes(t time.Time) bool {
	return !(t.Before(r.startTime) || t.After(r.endTime))
}

func (r TimeRange) eachDate() (result []time.Time) {
	for t := r.startTime; t.Before(r.endTime); t = t.AddDate(0, 0, 1) {
		result = append(result, t)
	}
	return
}

func YearRange(y int) TimeRange {
	startTime := time.Date(y, time.January, 1, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(y+1, time.January, 0, 0, 0, 0, 0, time.UTC)
	return TimeRange{startTime, endTime}
}

func MonthRange(m time.Month, y int) TimeRange {
	startTime := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(y, m+1, 0, 0, 0, 0, 0, time.UTC)
	return TimeRange{startTime, endTime}
}
