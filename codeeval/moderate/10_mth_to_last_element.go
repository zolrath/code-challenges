package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func mthToLast(list []string, m int) string {
	listlen := len(list)
	if m > listlen {
		return ""
	}
	return string(list[listlen-m])
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Split(s.Text(), " ")
		length := len(line)
		m, err := strconv.Atoi(line[length-1])
		if err != nil {
			log.Fatal(err)
		}
		line = line[:length-1]

		fmt.Println(mthToLast(line, m))
	}
}
