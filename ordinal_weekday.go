package recurrence

func OrdinalWeekday(w int, wd Weekday) Rule {
	return Intersection{Week(w), wd}
}
