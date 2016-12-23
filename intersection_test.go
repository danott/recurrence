package recurrence

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestIntersection(t *testing.T) {
	i := Intersection{
		January,
		Sunday,
	}
	r := YearRange(2006)

	assertIsOnlyOccurring(t, r, i, "2006-01-01", "2006-01-08", "2006-01-15",
		"2006-01-22", "2006-01-29")
}

func TestIntersectionOccurrences(t *testing.T) {
	tr := TimeRange{time.Time(NewDate("2006-01-01")), time.Time(NewDate("2009-12-31"))}

	expectations := map[int]Schedule{
		8: Intersection{Friday, Day(13)},
		4: Intersection{November, Thursday, Week(4)},
		3: Intersection{Day(Last), Day(28)},
	}

	assertOccurrenceGeneration2(t, tr, expectations)
}

func TestIntersectionMarshalJSON(t *testing.T) {
	tests := map[string]Intersection{
		`{"intersection":[{"day":1},{"day":"Last"},{"month":"January"}]}`:         Intersection{Day(First), Day(Last), January},
		`{"intersection":[{"weekday":"Thursday"},{"week":"Last"},{"year":2012}]}`: Intersection{Thursday, Week(Last), Year(2012)},
	}

	for expected, input := range tests {
		output, err := json.Marshal(input)
		if string(output) != expected || err != nil {
			t.Errorf("\nInput: %v\nExpected: %v\nActual: %v\nError: %v", input, expected, string(output), err)
		}
	}
}

func TestIntersectionUnmarshalJSON(t *testing.T) {
	tests := map[string]Intersection{
		`[{"day":"Last"},{"month":"January"}]`:   Intersection{Day(Last), January},
		`[{"weekday":"Thursday"},{"year":2014}]`: Intersection{Thursday, Year(2014)},
	}

	for input, expected := range tests {
		var output Intersection
		err := json.Unmarshal([]byte(input), &output)
		if !reflect.DeepEqual(output, expected) || err != nil {
			t.Errorf("\nInput: %v\nExpected: %v\nActual: %v\nError: %v", input, expected, output, err)
		}
	}
}

func TestIntersectionDayAndHourMinuteSecondOccuring(t *testing.T) {
	dhms := Intersection{Union{Saturday, Sunday}, NewHourMinuteSeconds(14, 30, 0)}
	// Today is Sunday 14/02/2016
	today := time.Date(2016, time.February, 14, 14, 30, 0, 0, time.UTC)

	if !dhms.IsOccurring(today) {
		t.Errorf("Event should be scheduled today: %v", today)
	}

	nextWeek := today.AddDate(0, 0, 6)
	if !dhms.IsOccurring(nextWeek) {
		t.Errorf("Event should be scheduled next Saturday too: %v", nextWeek)
	}

	nextWeek = nextWeek.AddDate(0, 0, 1)
	if !dhms.IsOccurring(nextWeek) {
		t.Errorf("Event should be scheduled next Sunday too: %v", nextWeek)
	}
}

func BenchmarkIntersectionOccurrences(b *testing.B) {
	d := Intersection{November, Thursday, Week(4)}
	tr := TimeRange{time.Now(), time.Now().AddDate(1000, 0, 0)}
	for n := 0; n < b.N; n++ {
		d.Occurrences(tr)
	}
}
