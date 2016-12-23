package recurrence

import (
	"testing"
	"time"
)

const f = "2006-01-02"

func assertIsOccurring(t *testing.T, s Schedule, o ...string) {
	for _, o := range o {
		if o, _ := time.Parse(f, o); !s.IsOccurring(o) {
			t.Errorf("%#v should have included %s", s, o)
		}
	}
}

func refuteIsOccurring(t *testing.T, s Schedule, o ...string) {
	for _, o := range o {
		if o, _ := time.Parse(f, o); s.IsOccurring(o) {
			t.Errorf("%#v should not have included %s", s, o)
		}
	}
}

func assertIsOnlyOccurring(t *testing.T, r TimeRange, s Schedule, o ...string) {
	asserted := make(map[string]bool)

	for _, o := range o {
		asserted[o] = true
	}

	for r := range r.eachDate() {
		if _, ok := asserted[r.Format(f)]; ok {
			assertIsOccurring(t, s, r.Format(f))
		} else {
			refuteIsOccurring(t, s, r.Format(f))
		}
	}
}

func assertAllOccurring(t *testing.T, r TimeRange, s Schedule) {
	for r := range r.eachDate() {
		assertIsOccurring(t, s, r.Format(f))
	}
}

func refuteAllOccurring(t *testing.T, r TimeRange, s Schedule) {
	for r := range r.eachDate() {
		refuteIsOccurring(t, s, r.Format(f))
	}
}

func assertOccurrenceGeneration(t *testing.T, tr TimeRange, expectations map[Schedule]int) {
	for schedule, expected := range expectations {
		schedule := schedule.(Schedule)
		var dates []time.Time

		for _, d := range schedule.Occurrences(tr) {
			dates = append(dates, d)
			if !schedule.IsOccurring(d) || d.Before(tr.Start) || d.After(tr.End) {
				t.Errorf("%s.Occurrences(%v) included a date it shouldn't have: %s", schedule, tr, d)
			}
		}

		if actual := len(dates); actual != expected {
			t.Errorf("%s.Occurrences should have generated %d. Got %d", schedule, expected, actual)
		}
	}
}

func assertOccurrenceGeneration2(t *testing.T, tr TimeRange, expectations map[int]Schedule) {
	for expected, schedule := range expectations {
		schedule := schedule.(Schedule)
		var dates []time.Time

		for _, d := range schedule.Occurrences(tr) {
			dates = append(dates, d)
			if !schedule.IsOccurring(d) || d.Before(tr.Start) || d.After(tr.End) {
				t.Errorf("%s.Occurrences(%v) included a date it shouldn't have: %s", schedule, tr, d)
			}
		}

		if actual := len(dates); actual != expected {
			t.Errorf("%s.Occurrences should have generated %d. Got %d", schedule, expected, actual)
		}
	}
}
