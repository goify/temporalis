package temporalis

import (
	"fmt"
	"time"
)

// pluralize returns the plural form of the given word if the count is not 1.
func pluralize(count int64, word string) string {
	if count == 1 {
		return fmt.Sprintf("%d %s", count, word)
	}
	return fmt.Sprintf("%d %ss", count, word)
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
