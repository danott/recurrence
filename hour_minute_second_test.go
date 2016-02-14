package recurrence

import (
	"testing"
	"time"
)

const hourLayout = "2006-01-02T15:04:05"

func assertIsHourOccurring(t *testing.T, s Schedule, o ...string) {
	for _, o := range o {
		if o, _ := time.Parse(hourLayout, o); !s.IsOccurring(o) {
			t.Errorf("%#v should have included %s", s, o)
		}
	}
}

func refuteIsHourOccurring(t *testing.T, s Schedule, o ...string) {
	for _, o := range o {
		if o, _ := time.Parse(hourLayout, o); s.IsOccurring(o) {
			t.Errorf("%#v should not have included %s", s, o)
		}
	}
}

func TestHourIsOccuring(t *testing.T) {
	hms := NewHourMinuteSeconds(14, 0, 0)
	assertIsHourOccurring(t, hms, "2016-01-01T14:00:00", "2016-01-02T14:00:00", "2016-01-20T14:00:00")

	hms = NewHourMinuteSeconds(14, 30, 0)
	assertIsHourOccurring(t, hms, "2016-01-01T14:30:00", "2016-01-02T14:30:00", "2016-01-20T14:30:00")

	hms = NewHourMinuteSeconds(14, 30, 59)
	assertIsHourOccurring(t, hms, "2016-01-01T14:30:59", "2016-01-02T14:30:59", "2016-01-20T14:30:59")
}

func TestHourIsNotOccuring(t *testing.T) {
	hms := NewHourMinuteSeconds(14, 0, 0)
	refuteIsHourOccurring(t, hms, "2016-01-01T14:01:00", "2016-01-04T20:00:00", "2016-01-20T07:00:00")

	hms = NewHourMinuteSeconds(14, 59, 0)
	refuteIsHourOccurring(t, hms, "2016-01-01T14:59:50", "2016-01-04T20:12:00", "2016-01-20T14:00:00")

	hms = NewHourMinuteSeconds(14, 0, 30)
	refuteIsHourOccurring(t, hms, "2016-01-01T14:00:29", "2016-01-04T20:00:30", "2016-01-20T14:00:00")
}

func TestNextHourMinuteSecondOccurrenceSameDay(t *testing.T) {
	hms := NewHourMinuteSeconds(14, 30, 0)
	d := time.Date(2006, time.January, 10, 14, 0, 0, 0, time.UTC)
	expected := time.Date(2006, time.January, 10, 14, 30, 0, 0, time.UTC)
	n, err := hms.nextAfter(d)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if expected != n {
		t.Errorf("Next occurences not matching [expected=%v, actual=%v]", expected, n)
	}
}

func TestNextHourMinuteSecondOccurrenceNextDay(t *testing.T) {
	hms := NewHourMinuteSeconds(14, 30, 0)
	d := time.Date(2006, time.January, 10, 16, 0, 0, 0, time.UTC)
	expected := time.Date(2006, time.January, 11, 14, 30, 0, 0, time.UTC)
	n, err := hms.nextAfter(d)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if expected != n {
		t.Errorf("Next occurences not matching [expected=%v, actual=%v]", expected, n)
	}
}

func TestNextHourMinuteSecondOccurences(t *testing.T) {
	hms := NewHourMinuteSeconds(14, 30, 0)
	tr := TimeRange{time.Time(NewDate("2006-01-01")), time.Time(NewDate("2006-01-31"))}
	nextDates := make([]time.Time, 0)

	for r := range tr.eachDate() {
		tn, err := hms.nextAfter(r)
		if err != nil {
			t.Errorf("Unable to get next time after [hms=%v, date=%v]: %v", hms, r, err)
		}
		nextDates = append(nextDates, tn)
	}

	expectedDays := 31
	if len(nextDates) != expectedDays {
		t.Errorf("Was expecting %d days but instead got %d instead", expectedDays, len(nextDates))
	}
}
