package recurrence

import (
	"encoding/json"
	"io"
	"time"
)

// The Schedule interface is the foundation of the recurrence package. Types
// satisfying the Schedule interface can be used to determine if a time.Time
// occurs in the Schedule, and generate time.Times satisfying the Schedule.
type Schedule interface {
	IsOccurring(time.Time) bool
	Occurrences(TimeRange) chan time.Time
}

func ScheduleFromJSON(r io.Reader) Schedule {
	var mixed interface{}
	var s Schedule
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&mixed); err != nil {
		panic("Decoding err'd")
	} else {
		scheduleKeyToValue(mixed, &s)
	}

	return s
}

// http://stackoverflow.com/questions/13364181/how-to-unmarshall-an-array-of-different-types-correctly
func scheduleKeyToValue(raw interface{}, schedule *Schedule) {
	for k, v := range raw.(map[string]interface{}) {
		switch k {
		case "Day":
			*schedule = Day(v.(float64))
		case "Weekday":
			*schedule = Weekday(v.(float64))
		case "Month":
			*schedule = Month(v.(float64))
		case "Week":
			*schedule = Week(v.(float64))
		case "Year":
			*schedule = Year(v.(float64))
		case "Intersection":
			var i Intersection
			for _, v := range v.([]interface{}) {
				var s Schedule
				scheduleKeyToValue(v, &s)
				i = append(i, s)
			}
			*schedule = i
		case "Union":
			var u Union
			for _, v := range v.([]interface{}) {
				var s Schedule
				scheduleKeyToValue(v, &s)
				u = append(u, s)
			}
			*schedule = u
		default:
			panic("!")
		}
	}
}
