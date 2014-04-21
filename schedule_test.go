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
