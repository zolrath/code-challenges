package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	stdin *bufio.Reader
)

func init() {
	stdin = bufio.NewReader(os.Stdin)
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func multigcd(nums []int) int {
	sortnums := make([]int, len(nums))
	copy(sortnums, nums)
	sort.Sort(sort.IntSlice(sortnums))
	m := sortnums[0]
	for i := 1; i < len(sortnums)-1; i++ {
		m = gcd(m, sortnums[i])
	}
	return m
}

func reduce(nums []int, gcd int) []int {
	for i, _ := range nums {
		nums[i] = nums[i] / gcd
	}
	return nums
}

func readInt() int {
	var n int
	for i, _ := stdin.ReadByte(); i >= '0' && i <= '9'; i, _ = stdin.ReadByte() {
		n = n*10 + int(i-'0')
	}
	return n
}

func printSlice(slice []int) {
	for i, v := range slice {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%d", v)
	}
	fmt.Print("\n")
}

func main() {
	var t, n int
	t = readInt()
	for i := 0; i < t; i++ {
		n = readInt()
		recipe := make([]int, 0, n)
		for j := 0; j < n; j++ {
			recipe = append(recipe, readInt())
		}
		gcd := multigcd(recipe)
		recipe = reduce(recipe, gcd)
		printSlice(recipe)
	}
}
