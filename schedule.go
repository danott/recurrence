package recurrence

import (
	"encoding/json"
	"fmt"
	"time"
)

// The Schedule interface is the foundation of the recurrence package. Types
// satisfying the Schedule interface can be used to determine if a time.Time
// occurs in the Schedule, and generate time.Times satisfying the Schedule.
type Schedule interface {
	IsOccurring(time.Time) bool
	Occurrences(TimeRange) chan time.Time
}

func UnmarshalJSON(b []byte) (schedule Schedule) {
	var mixed interface{}

	if err := json.Unmarshal(b, &mixed); err != nil {
		panic("Schedule could not be unmarshalled from bytes")
	} else {
		schedule = UnmarshalRawMap(mixed)
	}

	return schedule
}

// http://stackoverflow.com/questions/13364181/how-to-unmarshall-an-array-of-different-types-correctly
func UnmarshalRawMap(raw interface{}) (schedule Schedule) {
	for k, v := range raw.(map[string]interface{}) {
		switch k {
		case "Day":
			schedule = Day(v.(float64))
		case "Weekday":
			schedule = Weekday(v.(float64))
		case "Month":
			schedule = Month(v.(float64))
		case "Week":
			schedule = Week(v.(float64))
		case "Year":
			schedule = Year(v.(float64))
		case "Intersection":
			var intersection Intersection
			for _, v := range v.([]interface{}) {
				var member Schedule
				member = UnmarshalRawMap(v)
				intersection = append(intersection, member)
			}
			schedule = intersection
		case "Union":
			var union Union
			for _, v := range v.([]interface{}) {
				var member Schedule
				member = UnmarshalRawMap(v)
				union = append(union, member)
			}
			schedule = union
		case "Exclusion":
			var exclusion Exclusion
			for k, v := range v.(map[string]interface{}) {
				switch k {
				case "Schedule":
					var included Schedule
					included = UnmarshalRawMap(v)
					exclusion.Schedule = included
				case "Exclude":
					var excluded Schedule
					excluded = UnmarshalRawMap(v)
					exclusion.Exclude = excluded
				default:
					panic("Exclusion only accepts keys 'Exclude' and 'Schedule'")
				}
				schedule = exclusion
			}
		default:
			panic(fmt.Sprintf("Unknown Schedule type: %v", k))
		}
	}
	return
}
