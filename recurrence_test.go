package recurrence

import (
	"testing"
	"time"
)

const f = "2006-01-02"

func assertIncludes(t *testing.T, r Rule, s ...string) {
	for _, s := range s {
		if s, _ := time.Parse(f, s); !r.Includes(s) {
			t.Errorf("%#v should have included %s", r, s)
		}
	}
}

func assertExcludes(t *testing.T, r Rule, s ...string) {
	for _, s := range s {
		if s, _ := time.Parse(f, s); r.Includes(s) {
			t.Errorf("%#v should have excluded %s", r, s)
		}
	}
}

func assertOnlyIncludes(t *testing.T, tr TimeRange, r Rule, s ...string) {
	include := make(map[string]bool)

	for _, s := range s {
		include[s] = true
	}

	for aDate := range tr.eachDate() {
		if _, ok := include[aDate.Format(f)]; ok {
			assertIncludes(t, r, aDate.Format(f))
		} else {
			assertExcludes(t, r, aDate.Format(f))
		}
	}
}
