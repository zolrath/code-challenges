package main

import (
	"fmt"
	"strconv"
	"strings"
)

func multiplesTable(n, max int) {
	nums := make([]string, max)
	for i := 1; i <= max; i++ {
		nums[i-1] = strconv.Itoa(n * i)
	}
	fmt.Println(strings.Trim(fmt.Sprintf("%4v", nums), "[]"))

	//fmt.Printf("%4v", strconv.Itoa(n*i))
	//fmt.Println()
}

func main() {
	for i := 1; i <= 12; i++ {
		multiplesTable(i, 12)
	}
}
