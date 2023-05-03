package temporalis

import (
	"fmt"
	"strings"
	"time"
)

// After waits for the duration to elapse and then sends the current time on the returned channel.
// The function returns a channel that will receive the current time after the specified duration has passed.
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

// FormatDuration formats a time.Duration value into a human-readable string.
// The string will list each unit of time in descending order of magnitude,
// and will use the singular or plural form of the unit name as appropriate.
func FormatDuration(duration time.Duration) string {
	seconds := int64(duration.Seconds())

	days := seconds / 86400
	seconds -= days * 86400

	hours := seconds / 3600
	seconds -= hours * 3600

	minutes := seconds / 60
	seconds -= minutes * 60

	var parts []string
	if days > 0 {
		parts = append(parts, pluralize(days, "day"))
	}
	if hours > 0 {
		parts = append(parts, pluralize(hours, "hour"))
	}
	if minutes > 0 {
		parts = append(parts, pluralize(minutes, "minute"))
	}
	if seconds > 0 {
		parts = append(parts, pluralize(seconds, "second"))
	}

	switch len(parts) {
	case 0:
		return "0 seconds"
	case 1:
		return parts[0]
	case 2:
		return fmt.Sprintf("%s and %s", parts[0], parts[1])
	default:
		last := parts[len(parts)-1]
		parts = parts[:len(parts)-1]

		return fmt.Sprintf("%s and %s", strings.Join(parts, ", "), last)
	}
}

func BusinessHours(from, to time.Time, holidays []time.Time) time.Duration {
	var total time.Duration

	for from.Before(to) {
		if from.Weekday() != time.Saturday && from.Weekday() != time.Sunday && !isHoliday(from, holidays) {
			total += time.Hour
		}
		from = from.Add(time.Hour)
	}

	return total
}

func BusinessDays(from, to time.Time, holidays []time.Time) int {
	var total int

	for d := from; !d.After(to); d = d.AddDate(0, 0, 1) {
		if d.Weekday() != time.Saturday && d.Weekday() != time.Sunday && !isHoliday(d, holidays) {
			total++
		}
	}

	return total
}

func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func TimeDifference(from, to time.Time) time.Duration {
	return to.Sub(from)
}

func FormatTime(t time.Time, format string) string {
	return t.Format(format)
}

func UnixTimestamp(t time.Time) int64 {
	return t.Unix()
}

func FromUnixTimestamp(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

func TimezoneOffset(tz string, t time.Time) (int, error) {
	loc, err := time.LoadLocation(tz)

	if err != nil {
		return 0, err
	}
	_, offset := t.In(loc).Zone()

	return offset, nil
}

func TimezoneAbbreviation(tz string) (string, error) {
	loc, err := time.LoadLocation(tz)

	if err != nil {
		return "", err
	}

	now := time.Now().In(loc)

	return now.Format("MST"), nil
}
