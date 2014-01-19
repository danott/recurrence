package main

import "time"

const (
	First  = 1
	Second = 2
	Third  = 3
	Fourth = 4
	Fifth  = 5
	Last   = -1
)

const (
	Sunday time.Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type TemporalExpression interface {
	Includes(time.Time) bool
}

type Intersection []TemporalExpression

func (i Intersection) Includes(t time.Time) bool {
	for _, te := range i {
		if te.Includes(t) == false {
			return false
		}
	}
	return true
}

type Union []TemporalExpression

func (u Union) Includes(t time.Time) bool {
	for _, te := range u {
		if te.Includes(t) {
			return true
		}
	}
	return false
}

type Difference struct {
	Include []TemporalExpression
	Exclude []TemporalExpression
}

func (d Difference) Includes(t time.Time) bool {
	for _, te := range d.Exclude {
		if te.Includes(t) {
			return false
		}
	}
	for _, te := range d.Include {
		if te.Includes(t) {
			return true
		}
	}
	return false
}

type DateRange struct {
	Start time.Time
	End   time.Time
}

func (d DateRange) Includes(t time.Time) bool {
	return !(t.Before(d.Start) || t.After(d.End))
}

func (d DateRange) EachDate() (result []time.Time) {
	for t := d.Start; t.Before(d.End); t = t.AddDate(0, 0, 1) {
		result = append(result, t)
	}
	return
}

type WeekdayTE struct {
	Weekday time.Weekday
}

func (d WeekdayTE) Includes(t time.Time) bool {
	return t.Weekday() == d.Weekday
}

type WeekdayInMonthTE struct {
	Week    int
	Weekday time.Weekday
}

func (d WeekdayInMonthTE) Includes(t time.Time) (r bool) {
	if d.Week > 0 {
		r = (d.Weekday == t.Weekday() && weekFromMonthStart(t) == d.Week)
	} else {
		r = (d.Weekday == t.Weekday() && weekFromMonthEnd(t) == d.Week)
	}
	return
}

func weekFromMonthStart(t time.Time) int {
	return ((t.Day() - 1) / 7) + 1
}

func daysIn(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func weekFromMonthEnd(t time.Time) int {
	d := daysIn(t.Month(), t.Year())
	return (((d - t.Day()) / 7) * -1) - 1
}

func main() {
}
