package main

import (
	"fmt"
	"math"
)

func nPrimes(n int) []int {
	// Start with 2 in the primelist so we can only search odd numbers.
	primes := []int{2}
Primesearch:
	for i := 3; len(primes) < n; i += 2 {
		sqrt := int(math.Sqrt(float64(i)))
		for _, v := range primes {
			switch {
			case v > sqrt:
				break
			case i%v == 0:
				continue Primesearch
			}
		}
		primes = append(primes, i)
	}
	return primes
}

func main() {
	primes := nPrimes(1000)
	sum := 0
	for _, v := range primes {
		sum += v
	}
	fmt.Println(sum)
}
