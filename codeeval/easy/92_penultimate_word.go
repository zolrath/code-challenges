package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Split(s.Text(), " ")
		fmt.Println(line[len(line)-2])
	}
}
