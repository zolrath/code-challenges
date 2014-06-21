package bob

import "strings"

// hasLetters returns false if input string contains no letters.
func hasLetters(input string) bool {
	letterIndex := strings.IndexFunc(input, func(in rune) bool { return 'A' <= in && in <= 'z' })
	return letterIndex != -1
}

// isYelling returns true if string is all capitalized and contains letters.
func isYelling(input string) bool {
	return strings.ToUpper(input) == input && hasLetters(input)
}

// isQuestion returns true if the string ends with a question mark (?)
func isQuestion(input string) bool {
	return strings.HasSuffix(input, "?")
}

// isSilent returns true if the string only contained spaces or was empty.
func isSilent(input string) bool {
	return strings.TrimSpace(input) == ""
}

// Hey returns Bob's response to a given input.
func Hey(input string) string {
	switch {
	case isSilent(input):
		return "Fine. Be that way!"
	case isYelling(input):
		return "Woah, chill out!"
	case isQuestion(input):
		return "Sure."
	default:
		return "Whatever."
	}
}
