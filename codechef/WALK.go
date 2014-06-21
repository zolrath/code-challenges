package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	stdin *bufio.Reader
)

func init() {
	stdin = bufio.NewReader(os.Stdin)
}

func readInt() int {
	var n int
	p := 1

	for i, _ := stdin.ReadByte(); i >= '0' && i <= '9'; i, _ = stdin.ReadByte() {
		if i == '-' {
			p = -1
			continue
		}
		n = n*10 + int(i-'0')
	}
	return p * n
}

func getVelocity(weight []int) int {
	var reqvel, vel int
	for i, w := range weight {
		reqvel = w + i
		if reqvel > vel {
			vel = reqvel
		}
	}
	return vel
}

func main() {
	var t, n, v int
	// get number of test castes
	t = readInt()
	for i := 0; i < t; i++ {
		// get number of data points
		n = readInt()
		weight := make([]int, n)
		for j := 0; j < n; j++ {
			weight[j] = readInt()
		}
		v = getVelocity(weight)
		fmt.Println(v)
	}
}
