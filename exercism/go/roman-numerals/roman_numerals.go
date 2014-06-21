package romannumerals

import "strings"

type RomanNumerals []RomanNumeral

type RomanNumeral struct {
	decimal int
	roman   string
}

// ToRomanNumeral converts a given integer into a string of Roman Numerals.
func ToRomanNumeral(n int) string {
	result := []string{}
	numerals := RomanNumerals{
		{1000, "M"}, {900, "CM"},
		{500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"},
		{50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"},
		{5, "V"}, {4, "IV"},
		{1, "I"},
	}

	for _, num := range numerals {
		for n >= num.decimal {
			result = append(result, num.roman)
			n -= num.decimal
		}
	}

	return strings.Join(result, "")
}
