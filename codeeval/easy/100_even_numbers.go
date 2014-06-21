package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isEven(n int) int {
	if n%2 == 0 {
		return 1
	}
	return 0
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		num, _ := strconv.Atoi(s.Text())
		fmt.Println(isEven(num))
	}
}
