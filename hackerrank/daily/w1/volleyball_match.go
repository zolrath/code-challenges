package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func pathsToScore(a, b int) int {

}

func main() {
	s := bufio.NewScanner(os.Stdin)
	defer s.Close()

	s.Scan()
	a := strconv.Atoi(s.Text())
	s.Scan()
	b := strconv.Atoi(s.Text())
	fmt.Println(pathsToScore(a, b))
}
