package recurrence

import (
	"encoding/json"
	"fmt"
	"time"
)

// Date is a specific day. Shorthand for Intersection{Year, Month, Day}.
type Date time.Time

// Implement Schedule interface.
func (self Date) IsOccurring(t time.Time) bool {
	return beginningOfDay(time.Time(self)).Equal(beginningOfDay(t))
}

// Implement Schedule interface.
func (self Date) Occurrences(tr TimeRange) chan time.Time {
	ch := make(chan time.Time)

	go func() {
		start := tr.Start.AddDate(0, 0, -1)
		end := tr.End
		for t, err := self.NextAfter(start); err == nil && !t.After(end); t, err = self.NextAfter(t) {
			if !t.After(end) {
				ch <- t
			}
		}
		close(ch)
	}()

	return ch
}

func (self Date) NextAfter(t time.Time) (time.Time, error) {
	if t.Before(time.Time(self)) {
		return time.Time(self), nil
	}

	var zeroDate time.Time
	return zeroDate, fmt.Errorf("No more occurrences")
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
