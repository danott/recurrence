package recurrence

import (
	"encoding/json"
	"time"
)

// Date is a specific day. Shorthand for Intersection{Year, Month, Day}.
type Date time.Time

// Implement Schedule interface.
func (self Date) IsOccurring(t time.Time) bool {
	return self.asIntersection().IsOccurring(t)
}

// Implement Schedule interface.
func (self Date) Occurrences(t TimeRange) chan time.Time {
	return self.asIntersection().Occurrences(t)
}

func (self Date) asIntersection() Intersection {
	return Intersection{
		Year(time.Time(self).Year()),
		Month(time.Time(self).Month()),
		Day(time.Time(self).Day()),
	}
}

// Implement json.Unmarshaler interface.
func (self *Date) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(`"2006-01-02"`, string(b))

	if err != nil {
		return err
	}

	*self = Date(t)
	return nil
}

// Implement json.Marshaler interface.
func (self Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"date": time.Time(self).Format("2006-01-02"),
	})
}

// NewDate let's you create a new Date from the time format "2006-01-02"
func NewDate(s string) Date {
	t, err := time.Parse("2006-01-02", s)

	if err != nil {
		panic("NewDate requires format '2006-01-02'")
	}

	return Date(t)
}
