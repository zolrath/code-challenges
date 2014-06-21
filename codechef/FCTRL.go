package main

import (
	"bufio"
	"fmt"
	"os"
)

var in *bufio.Reader

func readInt() int {
	var result int
	for n, _ := in.ReadByte(); n >= '0' && n <= '9'; n, _ = in.ReadByte() {
		result = result*10 + int(n-'0')
	}
	return result
}

func powInt(num, pow int) int {
	result := num
	for ; pow > 1; pow-- {
		result = result * num
	}
	return result
}

func trailingZeroes(n int) int {
	z, pow := 0, 1
	for {
		f := n / powInt(5, pow)
		if f < 1 {
			break
		}
		z = z + f
		pow++
	}
	return z
}

func main() {
	in = bufio.NewReader(os.Stdin)
	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		fmt.Println(trailingZeroes(n))
	}
}
