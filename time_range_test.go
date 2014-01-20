package recurrence

import (
	"testing"
	"time"
)

func assertStartTimeAndEndTime(t *testing.T, r TimeRange, startTime string, endTime string) {
	expectedStartTime, _ := time.Parse(f, startTime)
	expectedEndTime, _ := time.Parse(f, endTime)

	if !r.startTime.Equal(expectedStartTime) {
		t.Errorf("%s should have been %s", r.startTime, expectedStartTime)
	}

	if !r.endTime.Equal(expectedEndTime) {
		t.Errorf("%s should have been %s", r.endTime, expectedEndTime)
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
