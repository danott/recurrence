package recurrence

import (
	"reflect"
	"testing"
)

func TestAnyScheduleUnmarshalJSON(t *testing.T) {
	tests := map[string]Schedule{
		`{"date":"2525-01-01"}`: NewDate("2525-01-01"),
		`{"day":"Last"}`:        Day(Last),
		`{"intersection":[{"day":1},{"day":"Last"},{"month":"January"}]}`:  Intersection{Day(First), Day(Last), January},
		`{"month":"January"}`:                                              January,
		`{"union":[{"weekday":"Thursday"},{"week":"Last"},{"year":2012}]}`: Union{Thursday, Week(Last), Year(2012)},
		`{"week":5}`:             Week(5),
		`{"weekday":"Saturday"}`: Saturday,
		`{"year":2525}`:          Year(2525),
	}

	for input, expected := range tests {
		output, err := ScheduleUnmarshalJSON([]byte(input))
		if !reflect.DeepEqual(output, expected) || err != nil {
			t.Errorf("\nInput: %v\nExpected: %v\nActual: %v\nError: %v", input, expected, output, err)
		}
	}
}
