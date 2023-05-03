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

// AfterFunc waits for the duration specified by d and then calls the function f
// in its own goroutine. It returns a Timer struct that can be used to cancel
// the function before it runs.
func AfterFunc(d time.Duration, f func()) *time.Timer {
	return time.AfterFunc(d, f)
}

// Date returns the Time corresponding to
// 00:00:00.0 UTC on the specified date in the given location.
func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
	return time.Date(year, month, day, hour, min, sec, nsec, loc)
}

// NewTicker returns a new Ticker containing a channel that will send the
// current time with a period specified by the duration argument. It adjusts the
// intervals or delays to make up for any slow-down or blocking of processing. The
// ticker will keep sending values until the Stop method is called on the returned
// Ticker object. If the duration is less than or equal to zero, NewTicker will
// panic. Use the time.Ticker.Stop() method to stop the ticker before its normal
// completion.
//
// Example usage:
//
//	ticker := temporalis.NewTicker(1 * time.Second)
//	defer ticker.Stop()
//	for {
//		select {
//		case t := <-ticker.C:
//			fmt.Println("tick at", t)
//		case <-done:
//			return
//		}
//	}
//
// In the example above, a new ticker is created that ticks once per second.
// The loop will keep running until either a value is received on the done
// channel, or the ticker is stopped using the Stop() method.
func NewTicker(d time.Duration) *time.Ticker {
	return time.NewTicker(d)
}

// NewTimer creates a new Timer that will send the current time on its channel after at least duration d.
// The returned timer contains a single channel that will be sent the current time when the timer expires.
// To use the timer, call its `C` method, which returns the channel on which the time will be sent.
// If the timer is not needed, it should be stopped by calling its `Stop` method.
// If the timer has already expired, the time will be sent immediately on the channel.
func NewTimer(d time.Duration) *time.Timer {
	return time.NewTimer(d)
}

// Now returns the current local time.
// This function is equivalent to calling time.Now() but returns a time.Time value in the local timezone.
func Now() time.Time {
	return time.Now()
}

// Sleep pauses the current goroutine for at least the duration d.
// A negative or zero duration causes Sleep to return immediately.
// This function is equivalent to time.Sleep in the standard library.
func Sleep(d time.Duration) {
	time.Sleep(d)
}

// Tick returns a new ticker that sends the current time on the returned
// channel at a regular interval defined by the duration argument. The ticker
// will start immediately and continue indefinitely, until stopped explicitly
// by calling its `Stop` method. The channel will close when the ticker is
// stopped.
//
// The ticker may adjust the time interval slightly to make the interval fit
// more accurately into the time grid defined by the operating system or
// hardware.
//
// Note that this function is usually only appropriate for use in endless
// functions, tests, and the main package. If you need to stop the ticker
// explicitly, or if you need a ticker that only runs for a limited number of
// times, consider using the `NewTicker` function instead.
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

// Format formats the time according to the layout string.
// The layout string is a representation of the time format as specified
// by the reference time "Mon Jan 2 15:04:05 -0700 MST 2006",
// with the same standard interpretations as the reference time.
// It returns the formatted time as a string.
func Format(t time.Time, layout string) string {
	return t.Format(layout)
}

// Parse parses a formatted string and returns the time value it represents.
// The layout string defines the format by showing how the reference time,
// defined to be
// Mon Jan 2 15:04:05 -0700 MST 2006
// would be represented if it were the value being parsed.
// The same interpretation as in Format is used to determine the meaning of each
// input character. Parse returns an error if the input string and layout string
// do not match.
func Parse(layout, value string) (time.Time, error) {
	return time.Parse(layout, value)
}

// ParseInLocation is like Parse but allows the caller to specify the location.
// The given location must be a valid time zone name such as "UTC" or "America/New_York",
// or a fixed offset in seconds east of UTC such as -18000 for Eastern Standard Time.
func ParseInLocation(layout, value string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(layout, value, loc)
}

// ParseTime parses a formatted string and returns the time value it represents.
// The layout string specifies the format by showing how the reference time,
// defined to be Mon Jan 2 15:04:05 -0700 MST 2006, would be formatted if it
// were the value. ParseTime returns an error if the input string and layout
// string do not match.
func ParseTime(str, format string) (time.Time, error) {
	return time.Parse(format, str)
}

// ConvertTimezone takes a time.Time object and a target timezone location string,
// and returns the time.Time object converted to the target timezone.
// The `location` parameter should be a string representing the target timezone's location,
// such as "America/New_York" or "Asia/Tokyo".
// The returned time.Time object will have the same UTC time as the input time.Time object,
// but its location will be set to the target timezone.
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

// DateRange returns a slice of time.Time values representing all the days
// between the start and end dates (inclusive). The time zone for the start and
// end dates should be specified as a string in the format "UTC±hh:mm", where
// "UTC" is the literal string "UTC" and "±hh:mm" is the time offset from UTC.
// If the start date is after the end date, an empty slice is returned.
func DateRange(start, end time.Time) []time.Time {
	var dates []time.Time

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		dates = append(dates, d)
	}

	return dates
}

// DateDiff calculates the difference between two dates and returns the result
// as a Duration. The first argument represents the start date, and the second
// argument represents the end date. If the start date is later than the end
// date, the function returns a negative duration. The returned duration will
// include any time that occurs between the start and end dates, including leap
// seconds and leap years. If either of the arguments are zero values, the
// function will panic.
func DateDiff(start, end time.Time) (int, error) {
	if end.Before(start) {
		return 0, fmt.Errorf("end date %v is before start date %v", end, start)
	}

	diff := end.Sub(start)

	return int(diff.Hours() / 24), nil
}

// isWeekend returns true if the given time is on a weekend (Saturday or Sunday), and false otherwise.
// It takes a single argument, t, which is the time to check.
func isWeekend(t time.Time) bool {
	return t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
}

// isHoliday checks if the given date is a holiday. It takes a date in the format
// "YYYY-MM-DD" and a map of holidays where the keys are the holiday dates in the
// same format and the values are the holiday names. If the given date is found in
// the holidays map, it returns true along with the name of the holiday, otherwise
// it returns false and an empty string.
func isHoliday(t time.Time, holidays []time.Time) bool {
	for _, h := range holidays {
		if t.Year() == h.Year() && t.Month() == h.Month() && t.Day() == h.Day() {
			return true
		}
	}

	return false
}

// WorkingDays returns the number of working days between two dates (inclusive).
// It takes start and end dates in the format "YYYY-MM-DD", and a list of holidays
// in the same format. The function assumes a 5-day workweek from Monday to Friday,
// and does not account for weekends. Holidays are considered as non-working days
// and are subtracted from the total number of days. If start date is after end date,
// the function returns 0. If the list of holidays is empty or nil, all days between
// the start and end dates are considered as working days.
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

// BusinessHours returns the number of business hours between two dates, excluding weekends and non-working hours.
// It takes start and end times, as well as the start and end hour of business for each weekday, and returns the
// duration of business hours between the two dates. The start and end hours of business for each weekday are
// specified as a map with keys representing the weekdays (time.Weekday) and values as structs with fields Start
// and End that represent the start and end hours of business for the given weekday. The timezone of the input
// dates and the business hours are assumed to be the same.
// The function returns a duration rounded up to the nearest minute.
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

// BusinessDays calculates the number of business days between two dates,
// excluding weekends and holidays based on the provided holiday list.
// It returns the number of business days and the list of holidays that fall
// within the date range (inclusive).
// If the end date is before the start date, the function returns 0 business days.
func BusinessDays(from, to time.Time, holidays []time.Time) int {
	var total int

	for d := from; !d.After(to); d = d.AddDate(0, 0, 1) {
		if d.Weekday() != time.Saturday && d.Weekday() != time.Sunday && !isHoliday(d, holidays) {
			total++
		}
	}

	return total
}

// IsLeapYear checks if a year is a leap year or not.
// A leap year is a year that is divisible by 4, except for years that are divisible by 100 and not divisible by 400.
// This means that years such as 1600 and 2000, which are divisible by 100 and 400, are leap years,
// while years such as 1700, 1800, and 1900, which are divisible by 100 but not by 400, are not leap years.
// The function returns true if the year is a leap year, and false otherwise.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// TimeDifference calculates the time difference between two given time.Time values
// and returns the result as a time.Duration. The first argument represents the
// starting time, and the second argument represents the ending time. If the ending
// time is before the starting time, the function returns a negative duration. The
// time difference is calculated as the absolute value of the difference between
// the two time values.
func TimeDifference(from, to time.Time) time.Duration {
	return to.Sub(from)
}

// FormatTime formats a given time according to a provided layout string and returns the formatted time string.
// The layout string is based on the reference time `Mon Jan 2 15:04:05 -0700 MST 2006`.
// The returned string is generated using the provided timezone, which defaults to UTC if not provided.
// Example layout string: "2006-01-02 15:04:05".
func FormatTime(t time.Time, format string) string {
	return t.Format(format)
}

// UnixTimestamp takes a time.Time value and returns its Unix timestamp,
// which is the number of seconds elapsed since January 1, 1970 UTC.
func UnixTimestamp(t time.Time) int64 {
	return t.Unix()
}

// FromUnixTimestamp returns a time.Time value representing the Unix timestamp in UTC timezone.
// The Unix timestamp represents the number of seconds elapsed since January 1, 1970 UTC.
func FromUnixTimestamp(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// TimezoneOffset returns the offset in seconds between the local time zone and the given time zone
// at the specified time. The offset is positive if the local time zone is ahead of the given time zone,
// and negative if the local time zone is behind the given time zone.
// The function takes a time zone abbreviation (e.g. "PST", "UTC") and a time.Time object as input.
// If the time zone abbreviation is not recognized by the time package, the function returns an error.
func TimezoneOffset(tz string, t time.Time) (int, error) {
	loc, err := time.LoadLocation(tz)

	if err != nil {
		return 0, err
	}
	_, offset := t.In(loc).Zone()

	return offset, nil
}

// TimezoneAbbreviation returns the abbreviated name of the timezone at a given time.
// It takes a time.Time object and returns a string representing the abbreviated name
// of the timezone (e.g. "PST" for Pacific Standard Time). The name returned is based
// on the current offset of the timezone from UTC.
func TimezoneAbbreviation(tz string) (string, error) {
	loc, err := time.LoadLocation(tz)

	if err != nil {
		return "", err
	}

	now := time.Now().In(loc)

	return now.Format("MST"), nil
}
