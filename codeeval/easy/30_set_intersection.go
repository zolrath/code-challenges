package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readSets(line string) ([]int, []int) {
	// Split line into strings of numbers
	setstxt := strings.Split(line, ";")
	setatxt := strings.Split(setstxt[0], ",")
	setbtxt := strings.Split(setstxt[1], ",")

	// Convert strings to ints for proper lexical comparisons.
	seta := make([]int, len(setatxt))
	setb := make([]int, len(setbtxt))
	for i, v := range setatxt {
		seta[i], _ = strconv.Atoi(v)
	}
	for i, v := range setbtxt {
		setb[i], _ = strconv.Atoi(v)
	}
	return seta, setb
}

// Creates intersection of two ascendingly sorted lists of ints.
func intersect(a, b []int) []int {
	union := []int{}
	alen, blen := len(a), len(b)
	i, j := 0, 0
	// Stop at end of both lists.
	for i < alen && j < blen {
		switch av, bv := a[i], b[j]; {
		// If values are the same, add to union.
		case av == bv:
			union = append(union, av)
			i++
			j++
		// If a is at a higher value than b, advance b.
		case av > bv:
			if j < blen {
				j++
			}
		// If b is at a higher value than a, advance a.
		case bv > av:
			if i < alen {
				i++
			}
		}
	}
	return union
}

func stringifySet(set []int) string {
	settxt := make([]string, len(set))
	for i, v := range set {
		settxt[i] = strconv.Itoa(v)
	}
	return strings.Join(settxt, ",")
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		seta, setb := readSets(s.Text())
		settxt := stringifySet(intersect(seta, setb))
		fmt.Println(settxt)
	}
}
