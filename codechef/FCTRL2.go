package main

import (
	"bufio"
	"fmt"
	"math/big"
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

func factorial(n int) *big.Int {
	r := big.NewInt(1)
	var i int
	for i = 1; i < n; i++ {
		r.Mul(r, big.NewInt((int64(i + 1))))
	}
	return r
}

func main() {
	var n, t int
	n = readInt()
	for i := 0; i < n; i++ {
		t = readInt()
		fmt.Println(factorial(t))
	}
}
