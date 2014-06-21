package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Split(s.Text(), " ")
		lineLen := len(line) - 1
		for i := 0; i <= (lineLen / 2); i++ {
			line[i], line[lineLen-i] = line[lineLen-i], line[i]
		}
		fmt.Println(strings.Join(line, " "))
	}
}
