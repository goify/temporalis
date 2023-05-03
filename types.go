package temporalis

type Duration int64

type Time struct {
	// wall and ext encode the wall time seconds, wall time nanoseconds,
	// and optional monotonic clock reading in nanoseconds.
	//
	// From high to low bit, wall encodes a 1-bit flag (hasMonotonic),
	// a 33-bit seconds field, and a 30-bit wall time nanoseconds field.
	// The nanoseconds field is in the range [0, 999999999].
	// If the hasMonotonic bit is 0, then the 33-bit field must be zero
	// and the full signed 64-bit wall seconds since Jan 1 year 1 is stored in ext.
	// If the hasMonotonic bit is 1, then the 33-bit field holds a 33-bit
	// unsigned wall seconds since Jan 1 year 1885, and ext holds a
	// signed 64-bit monotonic clock reading, nanoseconds since process start.
	wall uint64
	ext  int64
	loc  *Location
}

type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

var Months = [...]string{
	January:   "January",
	February:  "February",
	March:     "March",
	April:     "April",
	May:       "May",
	June:      "June",
	July:      "July",
	August:    "August",
	September: "September",
	October:   "October",
	November:  "November",
	December:  "December",
}

type Location struct {
	name string
	// zone specifies the set of rules to use in the current location.
	// The only zset variable value supported is "UTC",
	// for which the rules are hard-coded (in zoneinfo_unix.go).
	// All other values of zset are equivalent to "Local".
	zone []zone
	tx   []zoneTrans
}

type zone struct {
	name string
	// offset seconds east of UTC
	offset int
	// delta seconds to add to standard time to get wall clock time
	// aka Daylight Saving Time
	isDST bool
}

type zoneTrans struct {
	when         int64
	index        uint8
	isstd, isutc bool
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

var Weekdays = [...]string{
	Sunday:    "Sunday",
	Monday:    "Monday",
	Tuesday:   "Tuesday",
	Wednesday: "Wednesday",
	Thursday:  "Thursday",
	Friday:    "Friday",
	Saturday:  "Saturday",
}
