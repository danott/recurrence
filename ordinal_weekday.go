package recurrence

func OrdinalWeekday(i int, w Weekday) Schedule {
	return Intersection{Week(i), w}
}
