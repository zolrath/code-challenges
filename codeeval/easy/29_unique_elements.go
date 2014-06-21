package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Split(s.Text(), ",")
		uniq := map[int]bool{}

		// use map for uniqueness, convert to int for proper lexical sorting
		for _, v := range line {
			vi, _ := strconv.Atoi(v)
			uniq[vi] = true
		}

		// take map, put it into a slice for sorting
		list := []int{}
		for k, _ := range uniq {
			list = append(list, k)
		}
		sort.Ints(list)

		// convert back to string for join with ,
		stringlist := make([]string, len(list))
		for i, v := range list {
			v := strconv.Itoa(v)
			stringlist[i] = v
		}
		fmt.Println(strings.Join(stringlist, ","))

	}
}
