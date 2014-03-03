package recurrence

// Use OrdinalWeekday to generate Schedules for natural recurrence patterns such
// as the "Last Wednesday" or "Second Sunday".
func OrdinalWeekday(i int, w Weekday) Intersection {
	return Intersection{Week(i), w}
}
