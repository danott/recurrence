package recurrence

import "testing"

func TestWeek(t *testing.T) {
	tr := MonthRange(January, 2006)

	assertIsOnlyOccurring(t, tr, Week(First), "2006-01-01", "2006-01-02",
		"2006-01-03", "2006-01-04", "2006-01-05", "2006-01-06", "2006-01-07")

	assertIsOnlyOccurring(t, tr, Week(Second), "2006-01-08", "2006-01-09",
		"2006-01-10", "2006-01-11", "2006-01-12", "2006-01-13", "2006-01-14")

	assertIsOnlyOccurring(t, tr, Week(Third), "2006-01-15", "2006-01-16",
		"2006-01-17", "2006-01-18", "2006-01-19", "2006-01-20", "2006-01-21")

	assertIsOnlyOccurring(t, tr, Week(Fourth), "2006-01-22", "2006-01-23",
		"2006-01-24", "2006-01-25", "2006-01-26", "2006-01-27", "2006-01-28")

	assertIsOnlyOccurring(t, tr, Week(Fifth),
		"2006-01-29", "2006-01-30", "2006-01-31")

	assertIsOnlyOccurring(t, tr, Week(Last), "2006-01-25", "2006-01-26",
		"2006-01-27", "2006-01-28", "2006-01-29", "2006-01-30", "2006-01-31")
}
