package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isSquare(n int) bool {
	mod := n % 10
	if mod == 2 || mod == 3 || mod == 7 || mod == 8 {
		return false
	}
	rt := int(math.Sqrt(float64(n)) + 0.5)
	return rt*rt == n
}

func countSquares(start, end int) int {
	count := 0
	for i := start; i <= end; i++ {
		if isSquare(i) {
			count++
		}
	}
	return count
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	cases, _ := strconv.Atoi(s.Text())
	for i := 0; i < cases; i++ {
		s.Scan()
		numstr := strings.Split(s.Text(), " ")
		start, _ := strconv.Atoi(numstr[0])
		end, _ := strconv.Atoi(numstr[1])
		fmt.Println(countSquares(start, end))
	}
}
