package recurrence

import (
	"encoding/json"
	"testing"
)

func TestWeek(t *testing.T) {
	tr := MonthRange(January, 2006)

	assertIsOnlyOccurring(t, tr, Week(First), "2006-01-01", "2006-01-02",
		"2006-01-03", "2006-01-04", "2006-01-05", "2006-01-06", "2006-01-07")

	assertIsOnlyOccurring(t, tr, Week(Second), "2006-01-08", "2006-01-09",
		"2006-01-10", "2006-01-11", "2006-01-12", "2006-01-13", "2006-01-14")

	assertIsOnlyOccurring(t, tr, Week(Third), "2006-01-15", "2006-01-16",
		"2006-01-17", "2006-01-18", "2006-01-19", "2006-01-20", "2006-01-21")

	assertIsOnlyOccurring(t, tr, Week(Fourth), "2006-01-22", "2006-01-23",
		"2006-01-24", "2006-01-25", "2006-01-26", "2006-01-27", "2006-01-28")

	assertIsOnlyOccurring(t, tr, Week(Fifth),
		"2006-01-29", "2006-01-30", "2006-01-31")

	assertIsOnlyOccurring(t, tr, Week(Last), "2006-01-25", "2006-01-26",
		"2006-01-27", "2006-01-28", "2006-01-29", "2006-01-30", "2006-01-31")
}

func TestWeekMarshalJSON(t *testing.T) {
	tests := map[string]Week{
		`{"week":1}`:      Week(1),
		`{"week":2}`:      Week(2),
		`{"week":3}`:      Week(3),
		`{"week":4}`:      Week(4),
		`{"week":5}`:      Week(5),
		`{"week":"Last"}`: Week(Last),
	}

	for expected, input := range tests {
		output, err := json.Marshal(input)
		if string(output) != expected || err != nil {
			t.Errorf("\nInput: %v\nExpected: %v\nActual: %v\nError: %v", input, expected, output, err)
		}
	}
}

func TestWeekUnmarshalJSON(t *testing.T) {
	tests := map[string]Week{
		`1`:      Week(1),
		`2`:      Week(2),
		`3`:      Week(3),
		`4`:      Week(4),
		`5`:      Week(5),
		`"Last"`: Week(Last),
	}

	for input, expected := range tests {
		var output Week
		err := json.Unmarshal([]byte(input), &output)
		if output != expected || err != nil {
			t.Errorf("\nInput: %v\nExpected: %v\nActual: %v\nError: %v", input, expected, output, err)
		}
	}
}
