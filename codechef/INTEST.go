package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader *bufio.Reader
)

func readInt() int {
	var t int
	for i, _ := reader.ReadByte(); i >= '0' && i <= '9'; i, _ = reader.ReadByte() {
		t = t*10 + int(i-'0')
	}
	return t
}

func main() {
	var n, k, t, div int
	reader = bufio.NewReader(os.Stdin)
	n, k = readInt(), readInt()
	for i := 0; i < n; i++ {
		t = readInt()
		if t%k == 0 {
			div++
		}
	}
	fmt.Println(div)
}
