package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processLine(line string) (int, int) {
	split := strings.Split(line, ",")
	x, _ := strconv.Atoi(split[0])
	n, _ := strconv.Atoi(split[1])
	return x, n
}

// Not allowed to use division or modulo.
// Find smallest multiple of n greater than x.
func findMultiple(x, n int) int {
	multiple := n
	for multiple < x {
		multiple += n
	}
	return multiple
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		x, n := processLine(s.Text())
		fmt.Println(findMultiple(x, n))
	}
}
