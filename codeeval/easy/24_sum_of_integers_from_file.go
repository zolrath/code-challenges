package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()
	s := bufio.NewScanner(f)

	sum := 0
	for s.Scan() {
		line, _ := strconv.Atoi(s.Text())
		sum += line
	}
	fmt.Println(sum)
}
