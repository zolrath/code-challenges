package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
type Pair struct {
	Key   rune
	Value int
}

type PairList []Pair

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

func sortMapByValue(m map[rune]int) PairList {
	pl := make(PairList, len(m))
	i := 0
	for k, v := range m {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(pl)
	return pl
}*/

func scoreString(line string) int {
	line = strings.ToLower(line)
	freq := make([]int, 26)
	for _, v := range line {
		if 'a' <= v && v <= 'z' {
			freq[v-'a'] += 1
		}
	}
	sort.Ints(freq)
	value := 26
	score := 0
	for i := len(freq) - 1; i >= 0; i-- {
		score += freq[i] * value
		value--
	}
	return score
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		score := scoreString(s.Text())
		fmt.Println(score)
	}
}
