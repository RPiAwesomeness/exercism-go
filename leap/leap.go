/*
Package leap provides a simple function IsLeapYear which receives
an integer and returns a boolean depending on if it is a leap year.
*/

package leap

const testVersion = 3

// IsLeapYear takes an integer and checks if it's a leap year, divisible by 400
func IsLeapYear(year int) bool {
	if (year%4) == 0 && (year%100) != 0 {
		// Divisible by 4 but not 100, is a leap year
		return true
	} else if (year % 400) == 0 {
		// Divisible by 400, is a leap year
		return true
	} else if (year%100) == 0 && (year%400) != 0 {
		// Divisible by 100 but not 4, not a leap year
		return false
	}
	// Not divisible by 4, 100, or 400, not a leap year
	return false
}
