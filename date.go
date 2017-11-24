package recurrence

import (
	"encoding/json"
	"fmt"
	"time"
)

// Date is a specific day. Shorthand for Intersection{Year, Month, Day}.
type Date time.Time

// Implement Schedule interface.
func (d Date) IsOccurring(t time.Time) bool {
	return beginningOfDay(time.Time(d)).Equal(beginningOfDay(t))
}

// Implement Schedule interface.
func (d Date) Occurrences(tr TimeRange) chan time.Time {
	return occurrencesFor(d, tr)
}

func (d Date) nextAfter(t time.Time) (time.Time, error) {
	if t.Before(time.Time(d)) {
		return time.Time(d), nil
	}

	var zeroTime time.Time
	return zeroTime, fmt.Errorf("no more occurrences after %s", t)
}

// Implement json.Unmarshaler interface.
func (d *Date) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(`"2006-01-02"`, string(b))

	if err != nil {
		return err
	}

	*d = Date(t)
	return nil
}

// Implement json.Marshaler interface.
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"date": time.Time(d).Format("2006-01-02"),
	})
}

// NewDate let's you create a new Date from the time format "2006-01-02"
func NewDate(s string) Date {
	t, err := time.Parse("2006-01-02", s)

	if err != nil {
		panic(`NewDate(string) requires format "2006-01-02"`)
	}

	return Date(t)
}
