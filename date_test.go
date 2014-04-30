package recurrence

import (
	"encoding/json"
	"testing"
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
