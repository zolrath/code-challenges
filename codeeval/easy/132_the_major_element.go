package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func majorElement(line string) string {
	freq := map[string]int{}
	list := strings.Split(line, ",")
	limit := len(list) / 2
	for _, v := range list {
		freq[v] += 1
	}
	for k, v := range freq {
		if v >= limit {
			return k
		}
	}
	return "None"
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(majorElement(s.Text()))
	}
}
