package temporalis

import (
	"fmt"
	"time"
)

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
	ticker := time.NewTicker(d)
	done := make(chan struct{})
	c := make(chan time.Time)

	go func() {
		for {
			select {
			case t := <-ticker.C:
				c <- t
			case <-done:
				ticker.Stop()
				close(c)
				return
			}
		}
	}()

	return c
}

func Format(t time.Time, layout string) string {
	return t.Format(layout)
}

func Parse(layout, value string) (time.Time, error) {
	return time.Parse(layout, value)
}

func ParseInLocation(layout, value string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(layout, value, loc)
}

func ParseTime(str, format string) (time.Time, error) {
	return time.Parse(format, str)
}

func ConvertTimezone(t time.Time, from, to string) (time.Time, error) {
	locFrom, err := time.LoadLocation(from)

	if err != nil {
		return time.Time{}, err
	}

	locTo, err := time.LoadLocation(to)

	if err != nil {
		return time.Time{}, err
	}

	return t.In(locFrom).In(locTo), nil
}

func DateRange(start, end time.Time) []time.Time {
	var dates []time.Time

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		dates = append(dates, d)
	}

	return dates
}

func DateDiff(start, end time.Time) (int, error) {
	if end.Before(start) {
		return 0, fmt.Errorf("end date %v is before start date %v", end, start)
	}

	diff := end.Sub(start)

	return int(diff.Hours() / 24), nil
}

func isWeekend(t time.Time) bool {
	return t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
}

func isHoliday(t time.Time, holidays []time.Time) bool {
	for _, h := range holidays {
		if t.Year() == h.Year() && t.Month() == h.Month() && t.Day() == h.Day() {
			return true
		}
	}

	return false
}

func WorkingDays(start, end time.Time, holidays []time.Time) (int, error) {
	if end.Before(start) {
		return 0, fmt.Errorf("end date %v is before start date %v", end, start)
	}

	weekdays := 0

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		if !isWeekend(d) && !isHoliday(d, holidays) {
			weekdays++
		}
	}

	return weekdays, nil
}

func FormatDuration(duration time.Duration) string {
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	milliseconds := int(duration.Milliseconds()) % 1000

	if hours > 0 {
		return fmt.Sprintf("%d hours, %d minutes, %d seconds", hours, minutes, seconds)
	} else if minutes > 0 {
		return fmt.Sprintf("%d minutes, %d seconds", minutes, seconds)
	} else if seconds > 0 {
		return fmt.Sprintf("%d.%03d seconds", seconds, milliseconds)
	} else {
		return fmt.Sprintf("%d milliseconds", milliseconds)
	}
}
