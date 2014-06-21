package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func splitNum(n int) []int {
	nums := []int{}
	for n > 0 {
		nums = append(nums, n%10)
		n /= 10
	}
	return nums
}

func permutations(n int) [][]int {
	nums := splitNum(n)
	perms := [][]int{}
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			perms = append(perms, []int{nums[i], nums[j]})
		}
	}
	return perms
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	cases, _ := strconv.Atoi(s.Text())
	for i := 0; i < cases; i++ {
		s.Scan()
		line, _ := strconv.Atoi(s.Text())
		perms := permutations(line)
		fmt.Println(perms)
	}
}
