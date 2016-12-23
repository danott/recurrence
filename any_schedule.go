package recurrence

import (
	"encoding/json"
	"fmt"
	"time"
)

// AnySchedule acts as a wrapper around...any schedule. Why does this exists?
// Since the tree of Schedules is an arbitrary relationship of interfaces, we
// need a struct to *easily* marshal/unmarshal json of the entire schedule
// hierarchy.
type AnySchedule struct {
	Schedule Schedule
}

// Implement Schedule interface.
func (self AnySchedule) IsOccurring(t time.Time) bool {
	return self.Schedule.IsOccurring(t)
}

// Implement Schedule interface.
func (self AnySchedule) Occurrences(t TimeRange) []time.Time {
	return self.Schedule.Occurrences(t)
}

// Implement json.Marshaler interface.
func (d AnySchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Schedule)
}

// Implement json.Unmarshaler interface.
func (self *AnySchedule) UnmarshalJSON(b []byte) error {
	schedule, err := ScheduleUnmarshalJSON(b)

	if err != nil {
		return err
	} else {
		self.Schedule = schedule
	}

	return nil
}

// Unmarshal bytes representing any arbitrary relationship of schedules.
func ScheduleUnmarshalJSON(b []byte) (schedule Schedule, err error) {
	var mixed interface{}
	json.Unmarshal(b, &mixed)

	for key, value := range mixed.(map[string]interface{}) {
		rawValue, _ := json.Marshal(value)
		switch key {
		case "date":
			var date Date
			err = json.Unmarshal(rawValue, &date)
			schedule = date
		case "day":
			var day Day
			err = json.Unmarshal(rawValue, &day)
			schedule = day
		case "intersection":
			var intersection Intersection
			err = json.Unmarshal(rawValue, &intersection)
			schedule = intersection
		case "month":
			var month Month
			err = json.Unmarshal(rawValue, &month)
			schedule = month
		case "union":
			var union Union
			err = json.Unmarshal(rawValue, &union)
			schedule = union
		case "week":
			var week Week
			err = json.Unmarshal(rawValue, &week)
			schedule = week
		case "weekday":
			var weekday Weekday
			err = json.Unmarshal(rawValue, &weekday)
			schedule = weekday
		case "year":
			var year Year
			err = json.Unmarshal(rawValue, &year)
			schedule = year
		default:
			err = fmt.Errorf("%s is not a recognized schedule", key)
		}
	}
	return
}
