package atbash

import "unicode"

// encodeLetter encodes the letter with the Atbash cipher. a => z, b => y, z => a, y => b
func encodeLetter(input rune) rune {
	letter := unicode.ToLower(input)
	if 'a' <= letter && letter <= 'z' {
		return 'a' - letter + 'z'
	} else if '0' <= letter && letter <= '9' {
		return letter
	}
	return -1
}

// decodeLetter decodes the Atbash cipher letter to its original letter. z => a, y => b, a => z, b => y
func decodeLetter(input rune) rune {
	letter := unicode.ToLower(input)
	if 'a' <= letter && letter <= 'z' {
		return 'z' - letter + 'a'
	}
	return input
}

// Atbash encodes the given string with the Atbash cipher.
func Atbash(input string) string {
	result := make([]rune, 0, len(input))
	i := 0
	for _, l := range input {
		if encoded := encodeLetter(l); encoded != -1 {
			result = append(result, encoded)
			// If we've appended 5 letters, add a space and restart count.
			if i++; i == 5 {
				result = append(result, ' ')
				i = 0
			}
		}
	}

	// Trim space from end if present.
	l := len(result) - 1
	if result[l] == ' ' {
		result = result[:l]
	}

	return string(result)
}
