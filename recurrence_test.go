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

func assertOnlyIncludes(t *testing.T, r TimeRange, te Rule, s ...string) {
	include := make(map[string]bool)

	for _, s := range s {
		include[s] = true
	}

	for _, aDate := range r.eachDate() {
		if _, ok := include[aDate.Format(f)]; ok {
			assertIncludes(t, te, aDate.Format(f))
		} else {
			assertExcludes(t, te, aDate.Format(f))
		}
	}
}
