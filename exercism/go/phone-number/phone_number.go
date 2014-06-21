package phonenumber

import "fmt"

// isNumber returns true if rune is within '0' to '9'
func isNumber(input rune) bool {
	return '0' <= input && input <= '9'
}

// Number strips a given phone number of formatting.
func Number(num string) string {
	result := []rune{}
	for _, v := range num {
		if isNumber(v) {
			result = append(result, v)
		}
	}

	l := len(result)
	// If number is too large or small (allowing leading 1) return 0 string.
	if l < 10 || l > 11 || (l == 11 && result[0] != '1') {
		return "0000000000"
	}
	// If number has a leading 1 before area code, remove it.
	if l == 11 {
		return string(result[1:])
	}

	return string(result)
}

// AreaCode returns the area code of a given number.
func AreaCode(num string) string {
	return Number(num)[:3]
}

// Format returns a formatted string of a given number: (123) 456-7890
func Format(num string) string {
	n := Number(num)
	//return "(" + n[:3] + ") " + n[3:6] + "-" + n[6:] // Technically faster, technically ugly.
	return fmt.Sprintf("(%s) %s-%s", n[:3], n[3:6], n[6:])
}
