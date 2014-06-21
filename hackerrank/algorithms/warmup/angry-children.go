package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func minUnfairness(k int, packets []int) int {
	sort.Ints(packets)
	res := 100000000000
	for i := 0; i < len(packets)-k; i++ {
		diff := packets[i+k-1] - packets[i]
		if diff < res {
			res = diff
		}
	}
	return res
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	// Get number of candy packets
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	// Get number of children in village
	s.Scan()
	k, _ := strconv.Atoi(s.Text())

	packets := make([]int, n)

	for i := 0; i < n; i++ {
		s.Scan()
		packets[i], _ = strconv.Atoi(s.Text())
	}

	result := minUnfairness(k, packets)

	fmt.Println(result)
}
