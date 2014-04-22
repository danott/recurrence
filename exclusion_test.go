package recurrence

import "testing"

func TestExclusion(t *testing.T) {
	d := Exclusion{
		Day(Last),
		Union{September, November},
	}
	r := YearRange(2006)

	assertIsOnlyOccurring(t, r, d, "2006-01-31", "2006-02-28", "2006-03-31",
		"2006-04-30", "2006-05-31", "2006-06-30", "2006-07-31", "2006-08-31",
		"2006-10-31", "2006-12-31")
}
