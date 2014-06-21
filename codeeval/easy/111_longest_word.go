package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func longestWord(line string) string {
	var longestLen, longestIndex int
	words := strings.Split(line, " ")
	for i, v := range words {
		wordLen := len(v)
		if wordLen > longestLen {
			longestIndex = i
			longestLen = wordLen
		}
	}
	return words[longestIndex]
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		fmt.Println(longestWord(line))
	}
}
