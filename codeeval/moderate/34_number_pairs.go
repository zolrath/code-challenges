package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type NumSum struct {
	numbers []int
	sum     int
}

func (ns NumSum) GetPairs() [][]int {
	pairs := [][]int{}
	combos := combinations(ns.numbers)
	for _, v := range combos {
		total := 0
		total = v[0] + v[1]
		if total == ns.sum {
			pairs = append(pairs, v)
		}
	}
	return pairs
}

func intArrayToString(ints []int) []string {
	sa := make([]string, len(ints))
	for i, v := range ints {
		sa[i] = strconv.Itoa(v)
	}
	return sa
}

func (ns NumSum) String() string {
	pairs := ns.GetPairs()
	if len(pairs) == 0 {
		return "NULL"
	}
	var text string
	for _, v := range pairs {
		text += ";"
		sa := intArrayToString(v)
		text += strings.Join(sa, ",")
	}
	return text[1:]
}

func combinations(is []int) [][]int {
	combos := [][]int{}
	for i, _ := range is {
		for j := i + 1; j < len(is); j++ {
			combos = append(combos, []int{is[i], is[j]})
		}
	}
	return combos
}

func readFile(loc string) []NumSum {
	nums := []NumSum{}
	file, _ := os.Open(loc)
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := strings.Split(s.Text(), ";")
		pairs := strings.Split(line[0], ",")
		total, _ := strconv.Atoi(line[1])
		intpairs := []int{}
		for _, v := range pairs {
			iv, _ := strconv.Atoi(v)
			intpairs = append(intpairs, iv)
		}
		nums = append(nums, NumSum{intpairs, total})
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return nums
}

func main() {
	numbers := readFile(os.Args[1])
	for _, v := range numbers {
		fmt.Println(v)
	}
}
