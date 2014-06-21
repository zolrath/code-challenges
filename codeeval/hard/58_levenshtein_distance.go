package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("58.txt")
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
