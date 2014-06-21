package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func squareDigits(n int) int {
	sum := 0
	for n != 0 {
		digit := n % 10
		n /= 10
		sum += digit * digit
	}
	return sum
}

func isHappy(n int) int {
	chain := map[int]bool{}
	for n != 1 {
		n = squareDigits(n)
		if chain[n] == true {
			return 0
		}
		chain[n] = true
	}
	return 1
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		num, _ := strconv.Atoi(s.Text())
		fmt.Println(isHappy(num))
	}
}
