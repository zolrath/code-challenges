package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type IntReader struct {
	reader *bufio.Reader
}

func newIntReader(input io.Reader) *IntReader {
	reader := IntReader{bufio.NewReader(input)}
	return &reader
}

func (r *IntReader) ReadInt() int {
	var n int
	for i, _ := r.reader.ReadByte(); i >= '0' && i <= '9'; i, _ = r.reader.ReadByte() {
		n = n*10 + int(i-'0')
	}
	return n
}

func findSubsets(nums []int) [][]int {
	return [][]int{}
}

func xorSum(nums []int) int {
	return -1
}

func main() {
	r := newIntReader(os.Stdin)
	var t, n int
	t = r.ReadInt()
	// Iterate for t test cases.
	for i := 0; i < t; i++ {
		// n number of ints in test case.
		n = r.ReadInt()
		line := make([]int, n)
		// Accumulate all ints in slice of ints.
		for j := 0; j < n; j++ {
			line[j] = r.ReadInt()
		}
		fmt.Println(xorSum(line))
	}
}
