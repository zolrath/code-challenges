package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getSnakes(line string) map[int]int {
	snakes := make(map[int]int)
	split := strings.Split(line, " ")
	for _, p := range split {
		ohmy := strings.Split(p, ",")
		head, _ := strconv.Atoi(ohmy[0])
		tail, _ := strconv.Atoi(ohmy[1])
		snakes[head] = tail
	}
	return snakes
}

func getLadders(line string) map[int]int {
	ladders := make(map[int]int)
	split := strings.Split(line, " ")
	for _, p := range split {
		ohmy := strings.Split(p, ",")
		head, _ := strconv.Atoi(ohmy[0])
		tail, _ := strconv.Atoi(ohmy[1])
		ladders[head] = tail
	}
	return ladders
}

func hasLadder(n int, ladders map[int]int) int {
	furthestEnd := 0
	for i := 1; i <= 6; i++ {
		if end, ok := ladders[n+i]; ok {
			if end > furthestEnd {
				furthestEnd = end
			}
		}
	}
	return furthestEnd
}

func greatestNoSnakePos(n int, snakes map[int]int) int {
	closestSnakePos := 0

	for i := 6; i >= 1; i-- {
		if falldown, snake := snakes[n+i]; !snake {
			return n + i
		} else {
			if falldown > closestSnakePos {
				closestSnakePos = falldown
			}
		}
	}
	return closestSnakePos
}

func playGame(snakes, ladders map[int]int) int {
	pos := 1
	rolls := 0

	for pos < 100 {
		if newPos := hasLadder(pos, ladders); newPos > 0 {
			pos = newPos
			rolls++
		} else {
			pos = greatestNoSnakePos(pos, snakes)
			rolls++
		}
	}
	return rolls
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	s.Scan()
	testCases, _ := strconv.Atoi(s.Text())
	for i := 0; i < testCases; i++ {
		s.Scan()
		// Trash number of snakes/ladders
		s.Scan()
		ladders := getLadders(s.Text())
		s.Scan()
		snakes := getSnakes(s.Text())
		fmt.Println(playGame(snakes, ladders))
	}
}
