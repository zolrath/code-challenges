package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	X int
	Y int
}

func findDistance(a, b Point) int {
	xdist := float64(a.X - b.X)
	ydist := float64(a.Y - b.Y)
	return int(math.Sqrt((xdist * xdist) + (ydist * ydist)))
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		// Regexps are pretty gross. Maybe make a lexer for fun and profit?
		// Would definitely clean up a lot of this mess.
		re := regexp.MustCompile(`\((-?\d+), (-?\d+)\) \((-?\d+), (-?\d+)\)`)
		pointstring := re.FindStringSubmatch(line)
		points := make([]int, len(pointstring)-1)
		// Convert coordinants to ints, skipping 0 as it's the matched string.
		for i := 1; i < len(pointstring); i++ {
			points[i-1], _ = strconv.Atoi(pointstring[i])
		}
		xa, ya := points[0], points[1]
		xb, yb := points[2], points[3]
		pointa := Point{xa, ya}
		pointb := Point{xb, yb}
		fmt.Println(findDistance(pointa, pointb))
	}
}
