/*
An Armstrong number is an n-digit number that is equal to the sum of the n'th
powers of its digits. Determine if the input numbers are Armstrong numbers.
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func isArmstrong(num int) string {
	n := num
	digits := []int{}
	for n > 0 {
		digits = append(digits, (n % 10))
		n /= 10
	}
	p := float64(len(digits))
	sum := 0
	for _, v := range digits {
		sum += int(math.Pow(float64(v), p))
	}
	if num == sum {
		return "True"
	}
	return "False"
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		n, _ := strconv.Atoi(s.Text())
		fmt.Println(isArmstrong(n))
	}
}
