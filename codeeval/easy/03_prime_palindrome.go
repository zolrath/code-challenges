package main

import (
	"fmt"
	"math"
)

func primeSieve(limit int) []int {
	sieve := make([]int, limit+1)
	primes := []int{2}
	for i := 3; i <= int(math.Sqrt(float64(limit))); i += 2 {
		if sieve[i] == 0 {
			for j := i * i; j <= limit; j += i {
				sieve[j] = 1
			}
		}
	}
	for i := 3; i <= limit; i += 2 {
		if sieve[i] == 0 {
			primes = append(primes, i)
		}
	}
	return primes
}

func reverseInt(n int) int {
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n = n / 10
	}
	return reversed
}

func isIntPalindrome(n int) bool {
	return reverseInt(n) == n
}

func main() {
	primes := primeSieve(1000)
	for i := len(primes) - 1; i >= 0; i-- {
		if isIntPalindrome(primes[i]) {
			fmt.Println(primes[i])
			break
		}
	}
}
