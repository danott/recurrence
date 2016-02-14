package recurrence

import (
	"encoding/json"
	"testing"
	"time"
)

func TestWeekday(t *testing.T) {
	tr := MonthRange(January, 2006)

	assertIsOnlyOccurring(t, tr, Sunday,
		"2006-01-01", "2006-01-08", "2006-01-15", "2006-01-22", "2006-01-29")

	assertIsOnlyOccurring(t, tr, Monday,
		"2006-01-02", "2006-01-09", "2006-01-16", "2006-01-23", "2006-01-30")

	assertIsOnlyOccurring(t, tr, Tuesday,
		"2006-01-03", "2006-01-10", "2006-01-17", "2006-01-24", "2006-01-31")

	assertIsOnlyOccurring(t, tr, Wednesday,
		"2006-01-04", "2006-01-11", "2006-01-18", "2006-01-25")

	assertIsOnlyOccurring(t, tr, Thursday,
		"2006-01-05", "2006-01-12", "2006-01-19", "2006-01-26")

	assertIsOnlyOccurring(t, tr, Friday,
		"2006-01-06", "2006-01-13", "2006-01-20", "2006-01-27")

	assertIsOnlyOccurring(t, tr, Saturday,
		"2006-01-07", "2006-01-14", "2006-01-21", "2006-01-28")
}

func TestWeekdayOccurrences(t *testing.T) {
	tr := TimeRange{time.Time(NewDate("2006-01-01")), time.Time(NewDate("2006-12-31"))}

	expectations := map[Schedule]int{
		Sunday:    53,
		Monday:    52,
		Tuesday:   52,
		Wednesday: 52,
		Thursday:  52,
		Friday:    52,
		Saturday:  52,
	}

	assertOccurrenceGeneration(t, tr, expectations)
}

func TestWeekdayMarshalJSON(t *testing.T) {
	tests := map[string]Weekday{
		`{"weekday":"Sunday"}`:    Sunday,
		`{"weekday":"Monday"}`:    Monday,
		`{"weekday":"Tuesday"}`:   Tuesday,
		`{"weekday":"Wednesday"}`: Wednesday,
		`{"weekday":"Thursday"}`:  Thursday,
		`{"weekday":"Friday"}`:    Friday,
		`{"weekday":"Saturday"}`:  Saturday,
	}

	for expected, input := range tests {
		output, err := json.Marshal(input)
		if string(output) != expected || err != nil {
			t.Errorf("\nInput: %v\nExpected: %v\nActual: %v\nError: %v", input, expected, output, err)
		}
	}
}

func TestWeekdayUnmarshalJSON(t *testing.T) {
	tests := map[string]Weekday{
		`0`:           Sunday,
		`1`:           Monday,
		`2`:           Tuesday,
		`3`:           Wednesday,
		`4`:           Thursday,
		`5`:           Friday,
		`6`:           Saturday,
		`"Sunday"`:    Sunday,
		`"Monday"`:    Monday,
		`"Tuesday"`:   Tuesday,
		`"Wednesday"`: Wednesday,
		`"Thursday"`:  Thursday,
		`"Friday"`:    Friday,
		`"Saturday"`:  Saturday,
	}

	for input, expected := range tests {
		var output Weekday
		err := json.Unmarshal([]byte(input), &output)
		if output != expected || err != nil {
			t.Errorf("\nInput: %v\nExpected: %v\nActual: %v\nError: %v", input, expected, output, err)
		}
	}
}

func BenchmarkWeekdayOccurrences(b *testing.B) {
	d := Thursday
	tr := TimeRange{time.Now(), time.Now().AddDate(1000, 0, 0)}
	for n := 0; n < b.N; n++ {
		d.Occurrences(tr)
	}
}
