package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func rightmostChar(sent, let string) int {
	for i := len(sent) - 1; i >= 0; i-- {
		if string(sent[i]) == let {
			return i
		}
	}
	return -1
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Split(s.Text(), ",")
		sent, let := line[0], line[1]
		fmt.Println(rightmostChar(sent, let))
	}
}
