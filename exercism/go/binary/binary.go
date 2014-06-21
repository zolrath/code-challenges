package binary

// ToDecimal converts given binary string to its value as a decimal int.
func ToDecimal(binary string) int {
	result := 0
	binlen := len(binary) - 1

	for i := binlen; i >= 0; i-- {
		num := binary[i]
		if num == '1' {
			result += 1 << uint((binlen - i))
		} else if num != '0' {
			// If string contains something other than '0' or '1' return result 0.
			return 0
		}
	}
	return result
}
