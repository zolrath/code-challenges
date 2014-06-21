package main

import (
	"fmt"
	"strings"
)

// getInnermostBracket returns the indexes of the start and end point of the innermost set of brackets.
func getInnermostBracket(input string) (int, int) {
	lb := "({["
	rb := "]})"
	li := strings.LastIndexAny(input, lb)
	ri := strings.IndexAny(input[li:], rb) + li
	return li, ri
}

// debracket rips the string open from the inside out, putting the innermost bracket
// at the beginning of the result string and working outwards to each bracket pair.
func debracket(input string) string {
	result := make([]string, 0, len(input))
	for len(input) > 0 {
		left, right := getInnermostBracket(input)
		result = append(result, strings.Split(input[left+1:right], " ")...)
		input = input[:left] + input[right+1:]
	}
	return strings.Replace(strings.Join(result, " "), "  ", " ", -1)
}

func main() {
	input := []string{
		"((your[drink {remember to}]) ovaltine)",
		"[racket for {brackets (matching) is a} computers]",
		"[can {and it(it (mix) up ) } look silly]",
	}
	for _, v := range input {
		fmt.Println(debracket(v))
	}
}
