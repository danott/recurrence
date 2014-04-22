package recurrence

import "testing"

func TestUnion(t *testing.T) {
	u := Union{
		OrdinalWeekday(First, Sunday),
		Day(Last),
	}
	r := YearRange(2006)

	assertIsOnlyOccurring(t, r, u, "2006-01-01", "2006-02-05", "2006-03-05",
		"2006-04-02", "2006-05-07", "2006-06-04", "2006-07-02", "2006-08-06",
		"2006-09-03", "2006-10-01", "2006-11-05", "2006-12-03", "2006-01-31",
		"2006-02-28", "2006-03-31", "2006-04-30", "2006-05-31", "2006-06-30",
		"2006-07-31", "2006-08-31", "2006-09-30", "2006-10-31", "2006-11-30",
		"2006-12-31")
}

func TestIntersection(t *testing.T) {
	i := Intersection{
		January,
		Sunday,
	}
	r := YearRange(2006)

	assertIsOnlyOccurring(t, r, i, "2006-01-01", "2006-01-08", "2006-01-15",
		"2006-01-22", "2006-01-29")
}

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
