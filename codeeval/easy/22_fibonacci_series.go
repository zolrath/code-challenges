package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func nthFibonacci(n int) int {
	a, b := 1, 0
	for i := 0; i < n; i++ {
		a, b = a+b, a
	}
	return b
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		n, _ := strconv.Atoi(s.Text())
		fmt.Println(nthFibonacci(n))
	}
}
