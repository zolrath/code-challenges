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

func monsoonCycle(c int) int {
	height := 1
	for i := 1; i <= c; i++ {
		if i%2 != 0 {
			height *= 2
		} else {
			height += 1
		}
	}
	return height
}

func main() {
	var t int
	// get number of test castes
	t = readInt()
	for i := 0; i < t; i++ {
		fmt.Println(monsoonCycle(readInt()))
	}
}
