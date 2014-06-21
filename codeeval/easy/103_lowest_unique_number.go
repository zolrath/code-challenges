package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// LowestUnique returns the player with the lowest unique number, not the
// number itself. This is more compliated.
func lowestUnique(line []int) int {
	return 0
}

func main() {
	/*	f, _ := os.Open(os.Args[1])
		defer f.Close()*/

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := strings.Split(s.Text(), " ")
		list := make([]int, len(line))
		for i, v := range line {
			vi, _ := strconv.Atoi(v)
			list[i] = vi
		}
		fmt.Println(lowestUnique(list))
	}
}
