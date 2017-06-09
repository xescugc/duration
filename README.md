# duration

This package offers the hability to parse a string to a `time.Duration`. It's more strict (specific order and no repetition) than the [time.ParseDuration](parseDuration) but it offers more options: y (year) and w (week).

## Install

```
$> go get github.com/XescuGC/duration
```

## API

### ParseDuration(s string) (error, time.Duration)

ParseDuration parses a duration string with the following formats: `y`(year), `w`(week), `d`(day), `h`(hour), `m`(minute), `s`(second), `ms`(millisecond), `us` or `µs`(microsecond)  and `ns`(nanosecond). The order of the string is strict, must be from 'larger' to 'lesser', so '1m1s' is correct but '1s1m' is not, and repetitions are not permited, '1s1s' is invalid.

[parseDuration]: https://golang.org/pkg/time/#ParseDuration