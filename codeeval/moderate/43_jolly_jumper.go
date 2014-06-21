package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func isJollyJumper(seq []int) string {
	n := seq[0]
	seq = seq[1:]
	if n == 1 {
		return "Jolly"
	}
	var diff int
	diffs := make([]int, n)
	for i := 1; i < len(seq); i++ {
		diff = int(math.Abs(float64(seq[i-1] - seq[i])))
		if diff >= n || diffs[diff] == 1 {
			return "Not Jolly"
		}
		diffs[diff] = 1
	}
	for i := 1; i < len(diffs); i++ {
		if diffs[diff] == 0 {
			return "Not Jolly"
		}
	}
	return "Jolly"
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		lineints := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			lineints[i], err = strconv.Atoi(line[i])
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println(isJollyJumper(lineints))
	}
}
