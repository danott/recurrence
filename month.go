package recurrence

import "time"

type Month time.Month

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func (m Month) Includes(t time.Time) bool {
	return t.Month() == time.Month(m)
}
