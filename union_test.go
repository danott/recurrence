package recurrence

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestUnion(t *testing.T) {
	u := Union{
		OrdinalWeekday(First, Sunday),
		Day(Last),
	}
	r := YearRange(2006)

	assertIsOnlyOccurring(t, r, u, "2006-01-01", "2006-02-05", "2006-03-05",
		"2006-04-02", "2006-05-07", "2006-06-04", "2006-07-02", "2006-08-06",
		"2006-09-03", "2006-10-01", "2006-11-05", "2006-12-03", "2006-01-31",
		"2006-02-28", "2006-03-31", "2006-04-30", "2006-05-31", "2006-06-30",
		"2006-07-31", "2006-08-31", "2006-09-30", "2006-10-31", "2006-11-30",
		"2006-12-31")
}

func TestUnionOccurrences(t *testing.T) {
	tr := TimeRange{time.Time(NewDate("2006-01-01")), time.Time(NewDate("2009-12-31"))}

	expectations := map[int]Schedule{
		368: Union{June, July, August},
		626: Union{Monday, Wednesday, Friday},
		209: Union{Monday, Monday}, // Shouldn't duplicate days
	}

	assertOccurrenceGeneration2(t, tr, expectations)
}

func TestUnionMarshalJSON(t *testing.T) {
	tests := map[string]Union{
		`{"union":[{"day":1},{"day":"Last"},{"month":"January"}]}`:         Union{Day(First), Day(Last), January},
		`{"union":[{"weekday":"Thursday"},{"week":"Last"},{"year":2012}]}`: Union{Thursday, Week(Last), Year(2012)},
	}

	for expected, input := range tests {
		output, err := json.Marshal(input)
		if string(output) != expected || err != nil {
			t.Errorf("\nInput: %v\nExpected: %v\nActual: %v\nError: %v", input, expected, string(output), err)
		}
	}
}

func TestUnionUnmarshalJSON(t *testing.T) {
	tests := map[string]Union{
		`[{"day":"Last"},{"month":"January"}]`:   Union{Day(Last), January},
		`[{"weekday":"Thursday"},{"year":2014}]`: Union{Thursday, Year(2014)},
	}

	for input, expected := range tests {
		var output Union
		err := json.Unmarshal([]byte(input), &output)
		if !reflect.DeepEqual(output, expected) || err != nil {
			t.Errorf("\nInput: %v\nExpected: %v\nActual: %v\nError: %v", input, expected, output, err)
		}
	}
}

func BenchmarkUnionOccurrences(b *testing.B) {
	d := Union{January, March, May, Day(1), Day(Last)}
	tr := TimeRange{time.Now(), time.Now().AddDate(1000, 0, 0)}
	for n := 0; n < b.N; n++ {
		d.Occurrences(tr)
	}
}
