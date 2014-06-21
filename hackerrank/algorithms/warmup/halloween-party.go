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
	for i, _ := stdin.ReadByte(); i >= '0' && i <= '9'; i, _ = stdin.ReadByte() {
		n = n*10 + int(i-'0')
	}
	return n
}

func main() {
	var t, k int
	// get number of test castes
	t = readInt()
	for i := 0; i < t; i++ {
		// maximum number of cuts
		k = readInt()
		h := k / 2
		v := k - h
		fmt.Println(h * v)
	}
}
