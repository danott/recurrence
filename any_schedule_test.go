package recurrence

import "testing"

func TestAnyScheduleUnmarshalJSON(t *testing.T) {
	tests := map[string]Schedule{
		`{"month":"January"}`: January,
		`{"day":"Last"}`:      Day(Last),
		`{"year":2525}`:       Year(2525),
	}

	for input, expected := range tests {
		var output Schedule
		output, _ = ScheduleUnmarshalJSON([]byte(input))

		if output != expected {
			t.Errorf("\nInput: %#v\nExpected: %#v\nActual: %#v", input, expected, output)
		}
	}
}
