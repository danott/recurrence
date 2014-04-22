package recurrence

import "testing"

func TestIntersection(t *testing.T) {
	i := Intersection{
		January,
		Sunday,
	}
	r := YearRange(2006)

	assertIsOnlyOccurring(t, r, i, "2006-01-01", "2006-01-08", "2006-01-15",
		"2006-01-22", "2006-01-29")
}
