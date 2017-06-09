//This package offers the hability to parse a string to a time.Duration. Similar as the stdl but it offers two more options: y(year) and w(week).
package duration

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var (
	reDuration = regexp.MustCompile(`^(?:(?P<y>\d+)y)?(?:(?P<w>\d+)w)?(?:(?P<d>\d+)d)?(?:(?P<h>\d+)h)?(?:(?P<m>\d+)m)?(?:(?P<s>\d+)s)?(?:(?P<ms>\d+)ms)?(?:(?P<us>\d+)us)?(?:(?P<us>\d+)µs)?(?:(?P<ns>\d+)ns)?$`)
	reNames    = reDuration.SubexpNames()
	durations  = map[string]int64{
		"y":  int64(time.Hour) * 24 * 365,
		"w":  int64(time.Hour) * 24 * 7,
		"d":  int64(time.Hour) * 24,
		"h":  int64(time.Hour),
		"m":  int64(time.Minute),
		"s":  int64(time.Second),
		"ms": int64(time.Millisecond),
		"us": int64(time.Microsecond),
		"ns": int64(time.Nanosecond),
	}
)

// ParseDuration parses a duration string with the following formats: y(year), w(week), d(day), h(hour), m(minute), s(second), ms(millisecond), us or µs(microsecond)  and ns(nanosecond). The order of the string is strict, must be from 'larger' to 'lesser', so '1m1s' is correct but '1s1m' is not, and repetitions are not permited, '1s1s' is invalid.
func ParseDuration(s string) (time.Duration, error) {
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
