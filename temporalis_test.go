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
