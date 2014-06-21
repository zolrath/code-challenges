package bottles

import (
	"fmt"
	"strings"
)

// Maximum number of bottles permitted on wall as per Health Department regulation BW1099-4b.
const max = 99

// Verse returns the specific verse of a given number.
func Verse(n int) string {
	switch n {
	case 2:
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n"
	case 1:
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"
	case 0:
		return fmt.Sprintf("No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, %d bottles of beer on the wall.\n", max)
	default:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1)
	}
}

// Verses returns the verses of between the two given numbers, inclusive, as one string.
func Verses(upper, lower int) string {
	result := []string{}
	for i := upper; i >= lower; i-- {
		result = append(result, Verse(i))
	}
	// Append a blank string to ensure an extra new line at end of all verses.
	result = append(result, "")

	return strings.Join(result, "\n")
}

// Sing returns all the verses between 99 and 0.
func Sing() string {
	return Verses(max, 0)
}
