package temporalis

import "fmt"

// pluralize returns the plural form of the given word if the count is not 1.
func pluralize(count int64, word string) string {
	if count == 1 {
		return fmt.Sprintf("%d %s", count, word)
	}
	return fmt.Sprintf("%d %ss", count, word)
}
