package main

import (
	"bufio"
	"fmt"
	"os"
)

func swapCase(line string) string {
	newline := make([]rune, len(line))
	for i, v := range line {
		if 'a' <= v && v <= 'z' {
			newline[i] = v - 32
		} else if 'A' <= v && v <= 'Z' {
			newline[i] = v + 32
		} else {
			newline[i] = v
		}
	}
	return string(newline)
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		fmt.Println(swapCase(line))
	}
}
