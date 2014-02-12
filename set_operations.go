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
	Included Rule
	Excluded Rule
}

func (d Difference) Includes(t time.Time) bool {
	if d.Excluded.Includes(t) {
		return false
	}

	if d.Included.Includes(t) {
		return true
	}

	return false
}
