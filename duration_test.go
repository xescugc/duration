package duration

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	tests := []struct {
		Input    string
		Expected time.Duration
	}{
		{
			"1y",
			parseDuration("8760h"),
		},
		{
			"1w",
			parseDuration("168h"),
		},
		{
			"1d",
			parseDuration("24h"),
		},
		{
			"24h",
			parseDuration("24h"),
		},
		{
			"1m",
			parseDuration("1m"),
		},
		{
			"1s",
			parseDuration("1s"),
		},
		{
			"1ms",
			parseDuration("1ms"),
		},
		{
			"1us",
			parseDuration("1us"),
		},
		{
			"1µs",
			parseDuration("1µs"),
		},
		{
			"1ns",
			parseDuration("1ns"),
		},
		{
			"1d2h",
			parseDuration("26h"),
		},
	}

	for _, test := range tests {
		if d, _ := ParseDuration(test.Input); d != test.Expected {
			t.Errorf("Expected %s to be %d (%s), found %d (%s)", test.Input, test.Expected, test.Expected, d, d)
		}
	}

}

func TestParseDurationErrors(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{
			"as",
			"error parsing input: as",
		},
		{
			"1mas",
			"error parsing input: 1mas",
		},
		{
			"1s1m",
			"error parsing input: 1s1m",
		},
		{
			"1s1s",
			"error parsing input: 1s1s",
		},
	}

	for _, test := range tests {
		if _, err := ParseDuration(test.Input); err != nil {
			if err.Error() != test.Expected {
				t.Errorf("Expected %s to have error %s, found %q", test.Input, test.Expected, err)
			}
		} else {
			t.Errorf("Expected %s to have error %s, found %q", test.Input, test.Expected, err)
		}
	}
}

func parseDuration(s string) time.Duration {
	d, _ := time.ParseDuration(s)
	return d
}
