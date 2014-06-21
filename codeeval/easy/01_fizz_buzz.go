package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func fizzbuzz(a, b, n int) string {
	var fbstring string
	for i := 1; i <= n; i++ {
		fbstring += " "
		switch {
		case i%a == 0 && i%b == 0:
			fbstring += "FB"
		case i%a == 0:
			fbstring += "F"
		case i%b == 0:
			fbstring += "B"
		default:
			fbstring += strconv.Itoa(i)
		}
	}
	return fbstring[1:]
}

func readFile(loc string) [][]int {
	nums := [][]int{}
	file, _ := os.Open(loc)
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := strings.Split(s.Text(), " ")
		fbl := make([]int, 3)
		for i, v := range line {
			fbl[i], _ = strconv.Atoi(v)
		}
		nums = append(nums, fbl)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return nums
}
func main() {
	text := readFile(os.Args[1])
	for _, v := range text {
		fmt.Printf("%s\n", fizzbuzz(v[0], v[1], v[2]))
	}
}
