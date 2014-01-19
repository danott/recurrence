package recurrence

import "testing"

func TestWeekday(t *testing.T) {
	tr := MonthRange(January, 2006)

	assertOnlyIncludes(t, tr, Weekday{Sunday},
		"2006-01-01", "2006-01-08", "2006-01-15", "2006-01-22", "2006-01-29")

	assertOnlyIncludes(t, tr, Weekday{Monday},
		"2006-01-02", "2006-01-09", "2006-01-16", "2006-01-23", "2006-01-30")

	assertOnlyIncludes(t, tr, Weekday{Tuesday},
		"2006-01-03", "2006-01-10", "2006-01-17", "2006-01-24", "2006-01-31")

	assertOnlyIncludes(t, tr, Weekday{Wednesday},
		"2006-01-04", "2006-01-11", "2006-01-18", "2006-01-25")

	assertOnlyIncludes(t, tr, Weekday{Thursday},
		"2006-01-05", "2006-01-12", "2006-01-19", "2006-01-26")

	assertOnlyIncludes(t, tr, Weekday{Friday},
		"2006-01-06", "2006-01-13", "2006-01-20", "2006-01-27")

	assertOnlyIncludes(t, tr, Weekday{Saturday},
		"2006-01-07", "2006-01-14", "2006-01-21", "2006-01-28")
}
