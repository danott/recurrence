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
type Exclusion struct {
	Schedule Schedule
	Exclude  Schedule
}

func (d Exclusion) IsOccurring(t time.Time) bool {
	if d.Exclude.IsOccurring(t) {
		return false
	}

	if d.Schedule.IsOccurring(t) {
		return true
	}

	return false
}

func (d Exclusion) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(d)
}

func (d Exclusion) MarshalJSON() ([]byte, error) {
	type faux Exclusion
	return json.Marshal(struct {
		faux `json:"Exclusion"`
	}{faux: faux(d)})
}
