package recurrence

import "testing"

func TestAnyScheduleUnmarshalJSON(t *testing.T) {
	tests := map[string]Schedule{
		`{"month":"January"}`: January,
		`{"day":"Last"}`:      Day(Last),
		`{"year":2525}`:       Year(2525),
	}

	for input, expected := range tests {
		output, err := ScheduleUnmarshalJSON([]byte(input))
		if output != expected || err != nil {
			t.Errorf("\nInput: %v\nExpected: %v\nActual: %v\nError: %v", input, expected, output, err)
		}
	}
}
