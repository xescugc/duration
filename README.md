# duration

[![GoDoc](https://godoc.org/github.com/XescuGC/duration?status.svg)](https://godoc.org/github.com/XescuGC/duration)

This package offers the ability to parse a string to a `time.Duration`. It's more strict (specific order and no repetition) than the [time.ParseDuration](parseDuration) but it offers more options: y (year), w (week) and d(day).

It also has those durations as constants so: `duration.Year`, `duration.Week` and `duration.Day`

## Install

```
go get github.com/xescugc/duration
```

## API

### Parse(s string) (time.Duration, error)

Parse parses a duration string with the following formats:
* `y`(year)
* `w`(week)
* `d`(day)
* `h`(hour)
* `m`(minute)
* `s`(second)
* `ms`(millisecond)
* `us` or `µs`(microsecond)
* `ns`(nanosecond).

The order of the string is strict, must be from 'larger' to 'lesser', so `'1m1s'` is correct but `'1s1m'` is not, and repetitions are not permitted, `'1s1s'` is invalid.

[parseDuration]: https://golang.org/pkg/time/#ParseDuration
