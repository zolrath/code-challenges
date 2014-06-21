package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func modulo(n, m int) int {
	div := n / m
	mod := n - (m * div)
	return mod
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Split(s.Text(), ",")
		n, _ := strconv.Atoi(line[0])
		m, _ := strconv.Atoi(line[1])
		fmt.Println(modulo(n, m))
	}
}
