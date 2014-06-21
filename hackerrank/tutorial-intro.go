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

	// check if number is negative
	i, _ := stdin.ReadByte()
	if i == '-' {
		p = -1
		i = '0'
	}

	for ; i >= '0' && i <= '9'; i, _ = stdin.ReadByte() {
		n = n*10 + int(i-'0')
	}
	return p * n
}

func findElement(ar []int, e int) int {
	for i, v := range ar {
		if v == e {
			return i
		}
	}
	return -1
}

func main() {
	var v, n int
	// get value we're looking for
	v = readInt()
	// get number of elements in array
	n = readInt()
	// make array
	ar := make([]int, n)
	// populate array
	for i := 0; i < n; i++ {
		ar[i] = readInt()
	}
	fmt.Println(findElement(ar, v))
}
