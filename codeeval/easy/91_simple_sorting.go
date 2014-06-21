package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortLine(line []string) string {
	floatline := make([]float64, len(line))
	for i, v := range line {
		floatline[i], _ = strconv.ParseFloat(v, 64)
	}
	sort.Float64s(floatline)
	for i, v := range floatline {
		line[i] = strconv.FormatFloat(v, 'f', 3, 64)
	}
	return strings.Join(line, " ")
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Split(s.Text(), " ")
		fmt.Println(sortLine(line))
	}
}
