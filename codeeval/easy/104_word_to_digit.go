package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wtoi(word string) rune {
	ones := map[string]rune{
		"zero":  '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	return ones[word]
}

func wordstoi(words []string) string {
	digits := make([]rune, len(words))
	for i, v := range words {
		digits[i] = wtoi(v)
	}
	return string(digits)
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(wordstoi(strings.Split(s.Text(), ";")))
	}
}
