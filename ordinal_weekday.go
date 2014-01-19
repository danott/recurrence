package recurrence

import "time"

const (
	First  = 1
	Second = 2
	Third  = 3
	Fourth = 4
	Fifth  = 5
	Last   = -1
)

type OrdinalWeekday struct {
	Week    int
	Weekday time.Weekday
}

func (d OrdinalWeekday) Includes(t time.Time) (r bool) {
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
