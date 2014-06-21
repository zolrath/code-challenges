package octal

import "fmt"

// rtoi converts a rune '0'-'9' into an int64, returns an error if rune isn't a number.
func rtoi(n uint8) (int64, error) {
	if '0' <= n && n <= '9' {
		return int64(n - '0'), nil
	}
	return 0, fmt.Errorf("Rune %s is not a number.", string(n))
}

// ToDecimal converts a given octal string to a decimal int64.
func ToDecimal(oct string) int64 {
	var result, p int64
	p = 1
	for i := len(oct) - 1; i >= 0; i-- {
		dig, err := rtoi(oct[i])
		if err != nil {
			return 0
		}
		result += dig * p
		p *= 8
	}
	return result
}
