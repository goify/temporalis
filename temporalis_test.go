package temporalis

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	now1 := time.Now()
	now2 := Now()

	if now2.Sub(now1) > time.Second {
		t.Errorf("temporalis.Now() time difference too large: %v", now2.Sub(now1))
	}
}

func TestFormat(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	str := "2022-05-02 10:30:00"
	t1, _ := time.Parse(layout, str)
	formatted := Format(t1, layout)

	if formatted != str {
		t.Errorf("Formatted time incorrect: expected %q, but got %q", str, formatted)
	}
}

func TestParse(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	str := "2022-05-02 10:30:00"
	t1, _ := time.Parse(layout, str)
	parsed, err := Parse(layout, str)

	if err != nil {
		t.Errorf("Error parsing time: %v", err)
	}

	if parsed != t1 {
		t.Errorf("Parsed time incorrect: expected %v, but got %v", t1, parsed)
	}
}

// TestFormatDuration tests the FormatDuration function by providing several duration values and comparing the output with expected results.
// In this test, we check if the function formats a duration of 27 hours as "1 day and 3 hours" instead of the expected result "1 day, 2 hours and 1 minute".
// We also check if the function formats a duration of 28 hours as "1 day and 4 hours" instead of the expected result "1 day, 2 hours and 2 minutes".
func TestFormatDuration(t *testing.T) {
	tests := []struct {
		duration time.Duration
		expected string
	}{
		{0, "0 seconds"},
		{1 * time.Second, "1 second"},
		{2 * time.Second, "2 seconds"},
		{1 * time.Minute, "1 minute"},
		{2 * time.Minute, "2 minutes"},
		{1 * time.Hour, "1 hour"},
		{2 * time.Hour, "2 hours"},
		{24 * time.Hour, "1 day"},
		{48 * time.Hour, "2 days"},
		{25 * time.Hour, "1 day and 1 hour"},
		{26 * time.Hour, "1 day and 2 hours"},
		{27 * time.Hour, "1 day and 3 hours"},
		{28 * time.Hour, "1 day and 4 hours"},
		{2*24*time.Hour + 3*time.Hour + 4*time.Minute + 5*time.Second, "2 days, 3 hours, 4 minutes and 5 seconds"},
	}

	for _, test := range tests {
		actual := FormatDuration(test.duration)
		if actual != test.expected {
			t.Errorf("FormatDuration(%v) = %q, expected %q", test.duration, actual, test.expected)
		}
	}
}
