package recurrence

import (
	"encoding/json"
	"time"
)

// Computes the set intersection of a slice of Schedules.
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

func (i Intersection) MarshalJSON() ([]byte, error) {
	type faux Intersection
	return json.Marshal(struct {
		faux `json:"Intersection"`
	}{faux: faux(i)})
}

// Computes the set union of a slice of Schedules.
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

func (u Union) MarshalJSON() ([]byte, error) {
	type faux Union
	return json.Marshal(struct {
		faux `json:"Union"`
	}{faux: faux(u)})
}

// Computes the set difference of two Schedules.
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

func (d Difference) MarshalJSON() ([]byte, error) {
	type faux Difference
	return json.Marshal(struct {
		faux `json:"Difference"`
	}{faux: faux(d)})
}
