package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func isSelfDescribing(n []int) int {
	count := map[int]int{}
	for _, v := range n {
		count[v] += 1
	}
	for i, v := range n {
		if count[i] != v {
			return 0
		}
	}
	return 1
}

type IntReader struct {
	reader *bufio.Reader
}

func newIntReader(input io.Reader) *IntReader {
	reader := IntReader{bufio.NewReader(input)}
	return &reader
}

func (r *IntReader) ReadInt() []int {
	var n []int
	for i, _ := r.reader.ReadByte(); i >= '0' && i <= '9'; i, _ = r.reader.ReadByte() {
		n = append(n, int(i-'0'))
	}
	return n
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()
	r := newIntReader(f)
	for {
		nums := r.ReadInt()
		if nums == nil {
			break
		}
		fmt.Println(isSelfDescribing(nums))
	}
}
