package recurrence

import "testing"

func TestYear(t *testing.T) {
	y := Year(2006)

	refuteAllOccurring(t, YearRange(2005), y)
	assertAllOccurring(t, YearRange(2006), y)
	refuteAllOccurring(t, YearRange(2007), y)

	y = Year(2007)

	refuteAllOccurring(t, YearRange(2005), y)
	refuteAllOccurring(t, YearRange(2006), y)
	assertAllOccurring(t, YearRange(2007), y)
}
