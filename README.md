Recurrence
==========

Calendar recurrence rules in Go.

Basically, implementing the strategy described in [Martin Fowler's paper](http://martinfowler.com/apsupp/recurring.pdf). Read it. It's fun.

## Schedule

The `Schedule` interface is the foundation of the recurrence package. By using and combining schedules, we can represent all kinds of recurrence rules.

`schedule.IsOccurring(time.Time) bool` - does this `time.Time` occur in this schedule?
`schedule.Occurrences(TimeRange) chan time.Time` - generate `time.Time`s occuring with the passed `TimeRange`

## Day

Integer day of the month, 1 through 31, or the constant `Last`.

```go
first := Day(First) // Day(1)
last := Day(Last)
twentieth := Day(20)
```

## Week

Integer week of the month, 1 through 5, or the constant `Last`.

```go
first := Week(First) // Week(1)
last := Week(Last)
third := Week(Third) // Week(3)
```

## Weekday

Day of the week, Sunday through Monday. Constants are defined so you can use them with ease.

```go
Sunday.IsOccurring(time.Now())
```

## Month

A month of the year. Constants are defined to be used with ease.

```go
January.IsOccurring(time.Now())
```

## Set Operations

### Intersection

Intersection is a slice of Schedules. `IsOccurring` is only satisfied if all members of the slice are true. (Set intersection).

```go
// Complex Rules
american_thanksgiving := recurrence.Intersection{Week(4), Thursday, November}
```

### Union

Union is a slice of Schedules. `IsOccurring` is satisfied if any member of the slice is occurring. (Set union).

```go
weekends := recurrence.Union{Saturday, Sunday}
```

### Difference

Difference computes the set difference between two schedules.

```go
the_last_day_of_every_mont_except_september := recurrence.Difference{
   Day(Last),
   September,
}
```

