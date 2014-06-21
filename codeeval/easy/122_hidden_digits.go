package main

import (
	"bufio"
	"fmt"
	"os"
)

func findHidden(line string) string {
	hidden := []rune{}
	for _, v := range line {
		if 'a' <= v && v <= 'j' {
			hidden = append(hidden, v-'0'-1)
		} else if '0' <= v && v <= '9' {
			hidden = append(hidden, v)
		}
	}
	if len(hidden) == 0 {
		return "NONE"
	}
	return string(hidden)
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		fmt.Println(findHidden(line))
	}
}
