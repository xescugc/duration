// This package offers the hability to parse a string to a time.Duration, similar as the stdl but it offers three more options: y(year), w(week) and d(day).
package duration

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// This is the list of all the constant durations, it has the 3 added ones: Year, Week, Day and the others from
// the time.Duration so if using this package then it's not mixing both durations on use
const (
	Year        = time.Hour * 24 * 365
	Week        = time.Hour * 24 * 7
	Day         = time.Hour * 24
	Hour        = time.Hour
	Minute      = time.Minute
	Second      = time.Second
	Millisecond = time.Millisecond
	Microsecond = time.Microsecond
	Nanosecond  = time.Nanosecond
)

var (
	reDuration = regexp.MustCompile(`^(?:(?P<y>\d+)y)?(?:(?P<w>\d+)w)?(?:(?P<d>\d+)d)?(?:(?P<h>\d+)h)?(?:(?P<m>\d+)m)?(?:(?P<s>\d+)s)?(?:(?P<ms>\d+)ms)?(?:(?P<us>\d+)us)?(?:(?P<us>\d+)µs)?(?:(?P<ns>\d+)ns)?$`)
	reNames    = reDuration.SubexpNames()
	durations  = map[string]int64{
		"y":  int64(Year),
		"w":  int64(Week),
		"d":  int64(Day),
		"h":  int64(Hour),
		"m":  int64(Minute),
		"s":  int64(Second),
		"ms": int64(Millisecond),
		"us": int64(Microsecond),
		"ns": int64(Nanosecond),
	}
)

// ParseDuration parses a duration string with the following formats: y(year), w(week), d(day), h(hour), m(minute), s(second), ms(millisecond), us or µs(microsecond)  and ns(nanosecond). The order of the string is strict, must be from 'larger' to 'lesser', so '1m1s' is correct but '1s1m' is not, and repetitions are not permitted, '1s1s' is invalid.
func Parse(s string) (time.Duration, error) {
	res := reDuration.FindAllStringSubmatch(s, -1)
	var result int64

	for i, k := range reNames {
		if k == "" {
			continue
		}
		if len(res) == 0 {
			return 0, fmt.Errorf("error parsing input: %v", s)
		}
		if v := res[0][i]; v != "" {
			iv, err := strconv.Atoi(v)
			if err != nil {
				return 0, fmt.Errorf("error parsing input: %v", err)
			}
			result += durations[k] * int64(iv)
		}
	}

	return time.Duration(result), nil
}
