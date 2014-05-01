package recurrence

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDayIncludes(t *testing.T) {
	r := YearRange(2006)

	assertIsOnlyOccurring(t, r, Day(1), "2006-01-01", "2006-02-01", "2006-03-01",
		"2006-04-01", "2006-05-01", "2006-06-01", "2006-07-01", "2006-08-01",
		"2006-09-01", "2006-10-01", "2006-11-01", "2006-12-01")

	assertIsOnlyOccurring(t, r, Day(29), "2006-01-29", "2006-03-29", "2006-04-29",
		"2006-05-29", "2006-06-29", "2006-07-29", "2006-08-29", "2006-09-29",
		"2006-10-29", "2006-11-29", "2006-12-29")

	assertIsOnlyOccurring(t, r, Day(31), "2006-01-31", "2006-03-31", "2006-05-31",
		"2006-07-31", "2006-08-31", "2006-10-31", "2006-12-31")

	assertIsOnlyOccurring(t, r, Day(Last), "2006-01-31", "2006-02-28", "2006-03-31",
		"2006-04-30", "2006-05-31", "2006-06-30", "2006-07-31", "2006-08-31",
		"2006-09-30", "2006-10-31", "2006-11-30", "2006-12-31")
}

func TestDayOccurrences(t *testing.T) {
	tr := TimeRange{time.Time(NewDate("2006-01-01")), time.Time(NewDate("2009-12-31"))}

	expectations := map[Schedule]int{
		Day(1):    48,
		Day(2):    48,
		Day(3):    48,
		Day(4):    48,
		Day(5):    48,
		Day(27):   48,
		Day(28):   48,
		Day(29):   45,
		Day(30):   44,
		Day(31):   28,
		Day(Last): 48,
	}

	assertOccurrenceGeneration(t, tr, expectations)
}
func TestDayMarshalJSON(t *testing.T) {
	tests := map[string]Day{
		`{"day":1}`:      Day(1),
		`{"day":2}`:      Day(2),
		`{"day":3}`:      Day(3),
		`{"day":4}`:      Day(4),
		`{"day":5}`:      Day(5),
		`{"day":"Last"}`: Day(Last),
	}

	for expected, input := range tests {
		output, err := json.Marshal(input)
		if string(output) != expected || err != nil {
			t.Errorf("\nInput: %v\nExpected: %v\nActual: %v\nError: %v", input, expected, output, err)
		}
	}
}

func TestDayUnmarshalJSON(t *testing.T) {
	tests := map[string]Day{
		`1`:      Day(1),
		`2`:      Day(2),
		`3`:      Day(3),
		`4`:      Day(4),
		`5`:      Day(5),
		`6`:      Day(6),
		`7`:      Day(7),
		`8`:      Day(8),
		`9`:      Day(9),
		`10`:     Day(10),
		`11`:     Day(11),
		`12`:     Day(12),
		`13`:     Day(13),
		`14`:     Day(14),
		`15`:     Day(15),
		`16`:     Day(16),
		`17`:     Day(17),
		`18`:     Day(18),
		`19`:     Day(19),
		`20`:     Day(20),
		`21`:     Day(21),
		`22`:     Day(22),
		`23`:     Day(23),
		`24`:     Day(24),
		`25`:     Day(25),
		`26`:     Day(26),
		`27`:     Day(27),
		`28`:     Day(28),
		`29`:     Day(29),
		`30`:     Day(30),
		`31`:     Day(31),
		`"Last"`: Day(Last),
	}

	for input, expected := range tests {
		var output Day
		err := json.Unmarshal([]byte(input), &output)
		if output != expected || err != nil {
			t.Errorf("\nInput: %v\nExpected: %v\nActual: %v\nError: %v", input, expected, output, err)
		}
	}
}

func BenchmarkDayOccurrences(b *testing.B) {
	d := Day(1)
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
