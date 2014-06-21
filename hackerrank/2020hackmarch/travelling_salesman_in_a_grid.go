package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	cases, _ := strconv.Atoi(s.Text())
	for i := 0; i < cases; i++ {
		s.Scan()
		fmt.Println(s.Text())
	}
}
