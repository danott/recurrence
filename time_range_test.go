package recurrence

import (
	"testing"
	"time"
)

func assertStartTimeAndEndTime(t *testing.T, r TimeRange, start string, end string) {
	expectedStartTime, _ := time.Parse(f, start)
	expectedEndTime, _ := time.Parse(f, end)

	if !r.Start.Equal(expectedStartTime) {
		t.Errorf("%s should have been %s", r.Start, expectedStartTime)
	}

	if !r.End.Equal(expectedEndTime) {
		t.Errorf("%s should have been %s", r.End, expectedEndTime)
	}
}

func TestYearRange(t *testing.T) {
	assertStartTimeAndEndTime(t, YearRange(2006), "2006-01-01", "2006-12-31")
	assertStartTimeAndEndTime(t, YearRange(2014), "2014-01-01", "2014-12-31")
}

func TestMonthRange(t *testing.T) {
	assertStartTimeAndEndTime(t, MonthRange(time.January, 2006), "2006-01-01", "2006-01-31")
	assertStartTimeAndEndTime(t, MonthRange(time.February, 2006), "2006-02-01", "2006-02-28")
	assertStartTimeAndEndTime(t, MonthRange(time.February, 2008), "2008-02-01", "2008-02-29")
}

func TestMonthRangeIncludes(t *testing.T) {
	y := YearRange(2006)
	m := MonthRange(January, 2006)

	assertIsOnlyOccurring(t, y, m, "2006-01-01", "2006-01-02", "2006-01-03",
		"2006-01-04", "2006-01-05", "2006-01-06", "2006-01-07", "2006-01-08",
		"2006-01-09", "2006-01-10", "2006-01-11", "2006-01-12", "2006-01-13",
		"2006-01-14", "2006-01-15", "2006-01-16", "2006-01-17", "2006-01-18",
		"2006-01-19", "2006-01-20", "2006-01-21", "2006-01-22", "2006-01-23",
		"2006-01-24", "2006-01-25", "2006-01-26", "2006-01-27", "2006-01-28",
		"2006-01-29", "2006-01-30", "2006-01-31")
}

func TestMonthRangeAcceptableArguments(t *testing.T) {
	MonthRange(January, 2006)
	MonthRange(time.January, 2006)
	MonthRange(1, 2006)

	if x := recover(); x != nil {
		t.Errorf("Panicked on MonthRange tests")
	}
}
