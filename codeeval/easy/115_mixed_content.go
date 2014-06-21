package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func separate(content []string) string {
	words := []string{}
	nums := []string{}

	for _, elem := range content {
		n, err := strconv.Atoi(elem)
		if err != nil {
			words = append(words, elem)
		} else {
			nums = append(nums, strconv.Itoa(n))
		}
	}

	wordstring := strings.Join(words, ",")
	numstring := strings.Join(nums, ",")
	divider := ""
	if len(words) > 0 && len(nums) > 0 {
		divider = "|"
	}

	result := wordstring + divider + numstring
	return result
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(separate(strings.Split(s.Text(), ",")))
	}
}
