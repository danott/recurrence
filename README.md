Recurrence
==========

Calendar recurrence rules in Go.

Basically, implementing the strategy described in [Martin Fowler's paper](http://martinfowler.com/apsupp/recurring.pdf). Read it. It's fun.

Check out the [full docs on godoc.org](http://godoc.org/github.com/danott/recurrence)

## Schedule

The `Schedule` interface is the foundation of the recurrence package. By using and combining schedules, we can represent all kinds of recurrence rules.

`schedule.IsOccurring(time.Time) bool` - does this `time.Time` occur in this schedule?
`schedule.Occurrences(TimeRange) chan time.Time` - generate `time.Time`s occuring with the passed `TimeRange`

## Day

Integer day of the month, 1 through 31, or the constant `Last`.

```go
first := recurrence.Day(recurrence.First)
last := recurrence.Day(recurrence.Last)
twentieth := recurrence.Day(20)
```

## Week

Integer week of the month, 1 through 5, or the constant `Last`.

```go
first := recurrence.Week(recurrence.First)
last := recurrence.Week(recurrence.Last)
third := recurrence.Week(recurrence.Third)
```

## Weekday

Day of the week, Sunday through Monday. Constants are defined so you can use them with ease.

```go
recurrence.Sunday.IsOccurring(time.Now())
```

## Month

A month of the year. Constants are defined to be used with ease.

```go
recurrence.January.IsOccurring(time.Now())
```

## Year

A year.

```go
the_future := recurrence.Year(2525)
```

## TimeRange

A range of time. Primarily used as an argument to `schedule.Occurrences(t recurrence.TimeRange) chan time.Time`

It can also act as a schedule. Any day within the time range is considered as occurring.

```go
forty_days := recurrence.TimeRange{time.Now(), time.Now().AddDate(0, 0, 40)}
```

Some shortcuts are provided for common time ranges.

```go
recurrence.YearRange(2525)
recurrence.MonthRange(time.January, 2525)
```

## Set Operations

### Intersection

Intersection is a slice of Schedules. `IsOccurring` is only satisfied if all members of the slice are true. (Set intersection).

```go
american_thanksgiving := recurrence.Intersection{recurrence.Week(4), recurrence.Thursday, recurrence.November}
```

### Union

Union is a slice of Schedules. `IsOccurring` is satisfied if any member of the slice is occurring. (Set union).

```go
weekends := recurrence.Union{recurrence.Saturday, recurrence.Sunday}
```
