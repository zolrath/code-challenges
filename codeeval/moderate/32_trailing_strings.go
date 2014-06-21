package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isEnding(sentence, ending string) int {
	sentLen := len(sentence)
	endLen := len(ending)

	// If ending is longer then the sentence it can't be the ending.
	if endLen > sentLen {
		return 0
	}
	// start from end of string and compare.
	for i := 1; i <= endLen; i++ {
		if sentence[sentLen-i] != ending[endLen-i] {
			return 0
		}
	}
	return 1
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Split(s.Text(), ",")
		sentence := line[0]
		ending := line[1]
		fmt.Println(isEnding(sentence, ending))
	}
}
