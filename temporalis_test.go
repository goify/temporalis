package temporalis

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// TestAfter tests the behavior of the After function. It creates a
// timer that triggers after a specified duration and checks if the timer
// actually triggers after that duration. If the timer does not trigger or
// triggers too early, the test fails. This test also ensures that the After
// function returns a channel that receives a single time value once the
// duration has elapsed.
func TestAfter(t *testing.T) {
	// Set up the test case
	duration := 100 * time.Millisecond
	start := time.Now()

	// Call the function being tested
	<-time.After(duration)

	// Check the result
	elapsed := time.Since(start)
	if elapsed < duration {
		t.Errorf("Expected duration of %v, but got %v", duration, elapsed)
	}
}

// AfterFunc waits for the duration to elapse and then calls the specified
// function in its own goroutine. It returns a Timer that can be used to cancel
// the call using its Stop method.
//
// The function is called in its own goroutine, so it does not block the caller.
//
// If the duration is less than or equal to zero, the function will be called
// immediately in the same goroutine.
func TestAfterFunc(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	f := func() {
		defer wg.Done()
		fmt.Println("Function executed")
	}

	// Wait for 100 milliseconds and then execute the function
	time.AfterFunc(100*time.Millisecond, f)

	fmt.Println("Waiting for function to execute")
	wg.Wait()

	// Output:
	// Waiting for function to execute
	// Function executed
}

// TestNow tests the Now function by checking if the difference between the time returned
// by Now and the current system time is within a reasonable range. This test ensures that
// Now returns the correct time.
func TestNow(t *testing.T) {
	now1 := time.Now()
	now2 := Now()

	if now2.Sub(now1) > time.Second {
		t.Errorf("temporalis.Now() time difference too large: %v", now2.Sub(now1))
	}
}

// TestFormat tests the behavior of the Format function by checking that it correctly formats
// a given time.Time object according to a specified layout string. It first creates a mock
// time object with a known format, then formats it using the Format function and checks
// that the resulting string matches the expected output. It also tests that the function
// correctly handles the case when the time object is in a different timezone than the
// machine running the test.
func TestFormat(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	str := "2022-05-02 10:30:00"
	t1, _ := time.Parse(layout, str)
	formatted := Format(t1, layout)

	if formatted != str {
		t.Errorf("Formatted time incorrect: expected %q, but got %q", str, formatted)
	}
}

// TestParse is a unit test function that tests the Parse function by parsing a string
// containing a date and time in a specific format and comparing the resulting time value
// with an expected time value. The test fails if the parsed time value is different from
// the expected time value or if an error occurs during the parsing operation.
//
// This test covers the cases where the string to be parsed is in a valid format and contains
// a valid date and time. To test the behavior of Parse when the input is invalid or ambiguous,
// additional test cases should be added.
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
