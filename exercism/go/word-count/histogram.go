package wc

import "strings"

type Histogram map[string]int

func sanitize(input string) string {
	word := strings.Trim(input, "`'\"~!@#$%^&*()[]{}|\\/?-_=+:;.,")
	return strings.ToLower(word)
}

func WordCount(input string) Histogram {
	wc := Histogram{}
	words := strings.Split(input, " ")
	for _, word := range words {
		sw := sanitize(word)
		// Any individual punctuation mark sanitizes as an empty string, don't include.
		if len(sw) > 0 {
			wc[sw]++
		}
	}
	return wc
}
