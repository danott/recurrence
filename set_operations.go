package recurrence

import "time"

type Intersection []Rule

func (i Intersection) Includes(t time.Time) bool {
	for _, r := range i {
		if r.Includes(t) == false {
			return false
		}
	}
	return true
}

type Union []Rule

func (u Union) Includes(t time.Time) bool {
	for _, r := range u {
		if r.Includes(t) {
			return true
		}
	}
	return false
}

type Difference struct {
	Include []Rule
	Exclude []Rule
}

func (d Difference) Includes(t time.Time) bool {
	for _, r := range d.Exclude {
		if r.Includes(t) {
			return false
		}
	}
	for _, r := range d.Include {
		if r.Includes(t) {
			return true
		}
	}
	return false
}
