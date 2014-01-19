package recurrence

import "time"

const (
	January time.Month = 1 + iota
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

type Month struct {
	month time.Month
}

func (m Month) Includes(t time.Time) bool {
	return t.Month() == m.month
}
