package recurrence

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// AnySchedule acts as a wrapper around...any schedule. Why does this exists?
// Since the tree of Schedules is an arbitrary relationship of interfaces, we
// need a struct to *easily* marshal/unmarshal json of the entire schedule
// hierarchy.
type AnySchedule struct {
	Schedule `json:"schedule"`
}

func (self AnySchedule) IsOccurring(t time.Time) bool {
	return self.Schedule.IsOccurring(t)
}

func (self AnySchedule) Occurrences(t TimeRange) chan time.Time {
	return self.Schedule.Occurrences(t)
}

func (d AnySchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Schedule)
}

func (self *AnySchedule) UnmarshalJSON(b []byte) (err error) {
	var mixed interface{}

	err = json.Unmarshal(b, &mixed)
	if err != nil {
		return err
	} else {
		self.Schedule = unmarshalRawMap(mixed)
		if self.Schedule == nil {
			self.Schedule = Day(Never)
		}
	}

	return
}

// http://stackoverflow.com/questions/13364181/how-to-unmarshall-an-array-of-different-types-correctly
func unmarshalRawMap(raw interface{}) (schedule Schedule) {
	for k, v := range raw.(map[string]interface{}) {
		switch k {
		case "day":
			schedule = Day(v.(float64))
		case "weekday":
			schedule = Weekday(v.(float64))
		case "month":
			schedule = Month(v.(float64))
		case "week":
			schedule = Week(v.(float64))
		case "year":
			schedule = Year(v.(float64))
		case "intersection":
			var intersection Intersection
			for _, v := range v.([]interface{}) {
				var member Schedule
				member = unmarshalRawMap(v)
				intersection = append(intersection, member)
			}
			schedule = intersection
		case "union":
			var union Union
			for _, v := range v.([]interface{}) {
				var member Schedule
				member = unmarshalRawMap(v)
				union = append(union, member)
			}
			schedule = union
		case "exclusion":
			var exclusion Exclusion
			for k, v := range v.(map[string]interface{}) {
				switch k {
				case "schedule":
					var included Schedule
					included = unmarshalRawMap(v)
					exclusion.Schedule = included
				case "exclude":
					var excluded Schedule
					excluded = unmarshalRawMap(v)
					exclusion.Exclude = excluded
				default:
					panic("exclusion json only accepts 'exclude' and 'schedule' keys")
				}
			}
			if exclusion.Schedule == nil {
				panic(errors.New("exclusion json requires 'schedule' key"))
			}
			if exclusion.Exclude == nil {
				panic(errors.New("exclusion json requires 'exclude' key"))
			}
			schedule = exclusion
		default:
			panic(fmt.Sprintf("Unacceptable Schedule type: %v", k))
		}
	}

	if schedule == nil {
		schedule = Day(Never)
	}
	return
}
