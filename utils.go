package temporalis

// pluralize returns the plural suffix "s" if the given number is not 1.
func pluralize(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}
