package duration

import (
	"strconv"
	"strings"
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
		if d, _ := Parse(test.Input); d != test.Expected {
			t.Errorf("Expected %s to be %d (%s), found %d (%s)", test.Input, test.Expected, test.Expected, d, d)
		}
	}

}

func TestParseDurationLargeComponents(t *testing.T) {
	tests := []struct {
		Input    string
		Expected time.Duration
	}{
		{"2147483647ns", time.Duration(1<<31 - 1)},
		{"2147483648ns", time.Duration(1 << 31)},
		{"3000000000ns", 3 * time.Second},
		{"3000000000us", 50 * time.Minute},
		{"3000000000µs", 50 * time.Minute},
		{"3000000000ms", 3000000 * time.Second},
		{"3000000000s", 3000000000 * time.Second},
		{"1m3000000000ns", 63 * time.Second},
		{"9223372036854775807ns", time.Duration(1<<63 - 1)},
	}

	for _, test := range tests {
		t.Run(test.Input, func(t *testing.T) {
			got, err := Parse(test.Input)
			if err != nil {
				t.Fatalf("Parse(%q) returned an error: %v", test.Input, err)
			}
			if got != test.Expected {
				t.Errorf("Parse(%q) = %s, want %s", test.Input, got, test.Expected)
			}
		})
	}
}

func TestParseDurationComponentRangeErrors(t *testing.T) {
	for _, input := range []string{"9223372036854775808ns", "18446744073709551615ns"} {
		t.Run(input, func(t *testing.T) {
			if _, err := Parse(input); err == nil || !strings.Contains(err.Error(), strconv.ErrRange.Error()) {
				t.Errorf("Parse(%q) error = %v, want a range error", input, err)
			}
		})
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
		{
			"a",
			"error parsing input: a",
		},
		{
			"",
			"error parsing input: empty string",
		},
	}

	for _, test := range tests {
		if _, err := Parse(test.Input); err != nil {
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
