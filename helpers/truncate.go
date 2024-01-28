package helpers

// Truncate function truncates a string to the specified length.
func Truncate(s string, length int) string {
	if len(s) > length {
		return s[:length]
	}
	return s
}
