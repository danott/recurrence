package recurrence

import "time"

type Intersection []Schedule

func (i Intersection) IsOccurring(t time.Time) bool {
	for _, r := range i {
		if r.IsOccurring(t) == false {
			return false
		}
	}
	return true
}

func (i Intersection) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(i)
}

type Union []Schedule

func (u Union) IsOccurring(t time.Time) bool {
	for _, r := range u {
		if r.IsOccurring(t) {
			return true
		}
	}
	return false
}

func (u Union) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(u)
}

type Difference struct {
	Included Schedule
	Excluded Schedule
}

func (d Difference) IsOccurring(t time.Time) bool {
	if d.Excluded.IsOccurring(t) {
		return false
	}

	if d.Included.IsOccurring(t) {
		return true
	}

	return false
}

func (d Difference) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(d)
}
