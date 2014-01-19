package recurrence

import (
	"testing"
	"time"
)

func TestYearRange(t *testing.T) {
	r := YearRange(2006)
	expectedStartTime, _ := time.Parse(f, "2006-01-01")
	expectedEndTime, _ := time.Parse(f, "2006-12-31")
	if !r.startTime.Equal(expectedStartTime) {
		t.Errorf("%s should have been %s", r.startTime, expectedStartTime)
	}

	if !r.endTime.Equal(expectedEndTime) {
		t.Errorf("%s should have been %s", r.endTime, expectedEndTime)
	}
}

func TestMonthRange(t *testing.T) {
	r := MonthRange(time.January, 2006)
	expectedStartTime, _ := time.Parse(f, "2006-01-01")
	expectedEndTime, _ := time.Parse(f, "2006-01-31")
	if !r.startTime.Equal(expectedStartTime) {
		t.Errorf("%s should have been %s", r.startTime, expectedStartTime)
	}

	if !r.endTime.Equal(expectedEndTime) {
		t.Errorf("%s should have been %s", r.endTime, expectedEndTime)
	}
}
