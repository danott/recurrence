package recurrence

import (
	"encoding/json"
	"time"
)

type Date time.Time

func (self Date) IsOccurring(t time.Time) bool {
	return self.asIntersection().IsOccurring(t)
}

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

func (self *Date) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(`"2006-01-02"`, string(b))

	if err != nil {
		return err
	}

	*self = Date(t)
	return nil
}

func (self Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"date": time.Time(self).Format("2006-01-02"),
	})
}

func NewDate(s string) Date {
	t, err := time.Parse("2006-01-02", s)

	if err != nil {
		panic("NewDate requires format '2006-01-02'")
	}

	return Date(t)
}
