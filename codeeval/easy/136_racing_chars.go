package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	WALL       = '#'
	GATE       = '_'
	CHECKPOINT = 'C'
)

func findPath(paths string, entry int) int {
	likelyPath := -1
	left, right := 1, 1
	if entry-left < 0 {
		left = 0
	}
	if entry+right > len(paths) {
		right = 0
	}
	if entry == -1 {
		left = -1
		right = len(paths)
	}

	for i := entry - left; i <= entry+right; i++ {
		switch paths[i] {
		case WALL:
			continue
		case CHECKPOINT:
			return i
		case GATE:
			likelyPath = i
		}
	}
	return likelyPath
}

func carAngle(entry, exit int) rune {
	switch {
	case entry == -1:
		return '|'
	case entry > exit:
		return '/'
	case entry == exit:
		return '|'
	case entry < exit:
		return '\\'
	}
	return '|'
}

func drivePath(line string, entry int) (string, int) {
	path := findPath(line, entry)
	newLine := make([]rune, len(line))
	for i, v := range line {
		if i == path {
			newLine[i] = carAngle(entry, path)
			continue
		}
		newLine[i] = v
	}
	return string(newLine), path
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()
	s := bufio.NewScanner(f)
	entry := -1
	var newline string
	for s.Scan() {
		line := s.Text()
		newline, entry = drivePath(line, entry)
		fmt.Println(newline)
	}
}
