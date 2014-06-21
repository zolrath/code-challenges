package leap

// IsLeapYear returns true if the given year is a leap year, false if it's not.
// A leap year occurs on on every year that is evenly divisible by 4.
// Except every year that is evenly divisible by 100.
// Unless the year is also evenly divisible by 400.
func IsLeapYear(year int) bool {
	switch {
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	case year%4 == 0:
		return true
	default:
		return false
	}
}
