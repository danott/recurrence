package recurrence

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDateIncludes(t *testing.T) {
	r := YearRange(2006)

	assertIsOnlyOccurring(t, r, NewDate("2006-04-08"), "2006-04-08")
}

func TestDateMarshalJSON(t *testing.T) {
	tests := map[string]Date{
		`{"date":"2006-04-08"}`: NewDate("2006-04-08"),
		`{"date":"2525-01-01"}`: NewDate("2525-01-01"),
	}

	for expected, input := range tests {
		output, err := json.Marshal(input)
		if string(output) != expected || err != nil {
			t.Errorf("\nInput: %v\nExpected: %v\nActual: %v\nError: %v", input, expected, output, err)
		}
	}
}

func TestDateOccurrences(t *testing.T) {
	tr := TimeRange{time.Time(NewDate("2006-01-01")), time.Time(NewDate("2009-12-31"))}

	expectations := map[Schedule]int{
		NewDate("2005-12-31"): 0,
		NewDate("2006-01-01"): 1,
		NewDate("2007-04-08"): 1,
		NewDate("2009-12-31"): 1,
		NewDate("2010-01-01"): 0,
	}

	assertOccurrenceGeneration(t, tr, expectations)
}

func TestDateUnmarshalJSON(t *testing.T) {
	tests := map[string]Date{
		`"2006-04-08"`: NewDate("2006-04-08"),
		`"2525-01-01"`: NewDate("2525-01-01"),
	}

	for input, expected := range tests {
		var output Date
		err := json.Unmarshal([]byte(input), &output)
		if output != expected || err != nil {
			t.Errorf("\nInput: %v\nExpected: %v\nActual: %v\nError: %v", input, expected, output, err)
		}
	}
}

func BenchmarkDateOccurrences(b *testing.B) {
	d := NewDate("2525-01-01")
	tr := TimeRange{time.Now(), time.Now().AddDate(1000, 0, 0)}
	for n := 0; n < b.N; n++ {
		d.Occurrences(tr)
	}
}
