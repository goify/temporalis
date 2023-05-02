package temporalis

import (
	"time"
)

// Functions from the time package.

func After(d time.Duration) <-chan time.Time {
	return time.After(d)
}

func AfterFunc(d time.Duration, f func()) *time.Timer {
	return time.AfterFunc(d, f)
}

func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
	return time.Date(year, month, day, hour, min, sec, nsec, loc)
}

func NewTicker(d time.Duration) *time.Ticker {
	return time.NewTicker(d)
}

func NewTimer(d time.Duration) *time.Timer {
	return time.NewTimer(d)
}

func Now() time.Time {
	return time.Now()
}

func Sleep(d time.Duration) {
	time.Sleep(d)
}

func Tick(d time.Duration) <-chan time.Time {
	return time.Tick(d)
}

// Additional functions for formatting and parsing.

func Format(t time.Time, layout string) string {
	return t.Format(layout)
}

func Parse(layout, value string) (time.Time, error) {
	return time.Parse(layout, value)
}

func ParseInLocation(layout, value string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(layout, value, loc)
}
