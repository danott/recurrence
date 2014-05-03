package recurrence

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestYearIsOccurring(t *testing.T) {
	y := Year(2006)

	refuteAllOccurring(t, YearRange(2005), y)
	assertAllOccurring(t, YearRange(2006), y)
	refuteAllOccurring(t, YearRange(2007), y)

	y = Year(2007)

	refuteAllOccurring(t, YearRange(2005), y)
	refuteAllOccurring(t, YearRange(2006), y)
	assertAllOccurring(t, YearRange(2007), y)
}

func TestYearOccurrences(t *testing.T) {
	tr := NewTimeRange("2000-01-01", "3000-01-01")
	expectations := map[Schedule]int{
		Year(2525): 365,
	}
	assertOccurrenceGeneration(t, tr, expectations)

	tr = TimeRange{time.Time(NewDate("2525-12-31")), time.Time(NewDate("3000-01-01"))}
	expectations = map[Schedule]int{
		Year(2525): 1,
	}
	assertOccurrenceGeneration(t, tr, expectations)
}

func TestYearMarshalJSON(t *testing.T) {
	tests := map[string]Year{
		`{"year":2525}`: Year(2525),
	}

	// Get some arbitrary years in there.
	for i := 0; i < 10; i++ {
		random := rand.Intn(3535) + 1
		tests[fmt.Sprintf(`{"year":%d}`, random)] = Year(random)
	}

	for expected, input := range tests {
		output, err := json.Marshal(input)
		if string(output) != expected || err != nil {
			t.Errorf("\nInput: %v\nExpected: %v\nActual: %v\nError: %v", input, expected, output, err)
		}
	}
}

func TestYearUnmarshalJSON(t *testing.T) {
	tests := map[string]Year{
		`2525`: Year(2525),
	}

	// Get some arbitrary years in there.
	for i := 0; i < 10; i++ {
		random := rand.Intn(2233) + 1
		tests[fmt.Sprintf("%d", random)] = Year(random)
	}

	for input, expected := range tests {
		var output Year
		err := json.Unmarshal([]byte(input), &output)
		if output != expected || err != nil {
			t.Errorf("\nInput: %v\nExpected: %v\nActual: %v\nError: %v", input, expected, output, err)
		}
	}
}

func BenchmarkYearOccurrences(b *testing.B) {
	d := Year(2525)
	tr := TimeRange{time.Now(), time.Now().AddDate(1000, 0, 0)}
	for n := 0; n < b.N; n++ {
		ch := d.Occurrences(tr)
		for {
			_, ok := <-ch

			if !ok {
				break
			}
		}
	}
}
