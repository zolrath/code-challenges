package anagram

import (
	"sort"
	"strings"
)

// RuneSlice type allows us to satisfy the Sort interface with Swap, Len, and Less.
type RuneSlice []rune

func (rs RuneSlice) Swap(i, j int)      { rs[i], rs[j] = rs[j], rs[i] }
func (rs RuneSlice) Len() int           { return len(rs) }
func (rs RuneSlice) Less(i, j int) bool { return rs[i] < rs[j] }

// sortWord takes a word and sorts it by rune value.
func sortWord(word string) string {
	runes := make(RuneSlice, len(word))
	for i, v := range word {
		runes[i] = v
	}
	sort.Sort(runes)
	return string(runes)
}

// Operating on runes takes roughly 2/3rds as much time.
/*func sortWord(word string) string {
	letters := strings.Split(word, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}*/

// Detect takes a word and a slice of anagram candidates and returns a slice of
// actual anagrams in the candidate slice.
func Detect(word string, candidates []string) []string {
	anagrams := []string{}
	// Make word lowercase for normalization.
	word = strings.ToLower(word)

	// Sort the word for comparison to sorted candidate later.
	// Don't sort the original string so we can remove candidates that are the same as word.
	sortedWord := sortWord(word)

	for _, candidate := range candidates {
		// The candidate and word have to be the same length to be anagrams.
		if len(candidate) != len(word) {
			continue
		}

		// Make candidate lowercase for normalization.
		candidate = strings.ToLower(candidate)

		// If the candidate is the word it's not an anagram.
		if word == candidate {
			continue
		}

		// Sort the candidate for comparison to sorted word.
		sortedCandidate := sortWord(candidate)
		if sortedWord == sortedCandidate {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}
