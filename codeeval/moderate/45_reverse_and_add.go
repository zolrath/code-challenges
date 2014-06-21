package main

import (
	"fmt"
	"strconv"
)
import "log"
import "bufio"
import "os"

func reverseInt(n int) int {
	var reversed int
	for n > 0 {
		reversed = reversed*10 + n%10
		n = n / 10
	}
	return reversed
}

func isIntPalindrome(n int) bool {
	return n == reverseInt(n)
}

func sumWithReverse(n int) int {
	return n + reverseInt(n)
}

func sumUntilPalindrome(n int) (int, int) {
	var iter int
	for {
		if isIntPalindrome(n) {
			break
		}
		n = sumWithReverse(n)
		iter++
	}
	return iter, n
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		i, p := sumUntilPalindrome(n)
		fmt.Println(i, p)
	}
}
