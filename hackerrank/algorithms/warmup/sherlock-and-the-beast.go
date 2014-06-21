package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func decentNumber(n int) string {
	maxfives := n / 3
	maxthrees := n / 5
	for i := maxfives; i >= 0; i-- {
		for j := 0; j <= maxthrees; j++ {
			if 3*i+5*j == n {
				return strings.Repeat("5", i*3) + strings.Repeat("3", j*5)
			}
		}
	}
	return "-1"
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	t, _ := strconv.Atoi(s.Text())
	for i := 0; i < t; i++ {
		s.Scan()
		c, _ := strconv.Atoi(s.Text())
		fmt.Println(decentNumber(c))
	}
}
