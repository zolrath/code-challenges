package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	line := strings.Split(s.Text(), " ")
	// Get number of jars.
	n, _ := strconv.Atoi(line[0])
	// Get number of operations.
	m, _ := strconv.Atoi(line[1])

	totalCandies := 0

	// Perform each operation.
	for i := 0; i < m; i++ {
		s.Scan()
		line = strings.Split(s.Text(), " ")
		// Get lower range of jars.
		a, _ := strconv.Atoi(line[0])
		// Get upper range of jars.
		b, _ := strconv.Atoi(line[1])
		// Get candy amount to add.
		k, _ := strconv.Atoi(line[2])
		// Add candies to pile of candy.
		totalCandies += (b - a + 1) * k
	}

	averageCandies := totalCandies / n

	fmt.Println(averageCandies)
}
