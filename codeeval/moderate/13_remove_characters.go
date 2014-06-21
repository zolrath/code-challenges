package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func removeChars(str string, remove string) string {
	var newstring []rune
LetterSearch:
	for _, l := range str {
		for _, letter := range remove {
			if l == letter {
				continue LetterSearch
			}
		}
		newstring = append(newstring, l)
	}
	return string(newstring)
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		input := strings.Split(s.Text(), ",")
		text := input[0]
		remove := input[1][1:]
		fmt.Println(removeChars(text, remove))
	}
}
