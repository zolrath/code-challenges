package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processLine(line string) (int, int, int) {
	split := strings.Split(line, ",")
	num, _ := strconv.Atoi(split[0])
	bita, _ := strconv.Atoi(split[1])
	bitb, _ := strconv.Atoi(split[2])
	return num, bita, bitb
}

func compareBits(num int, posa, posb uint) bool {
	posa--
	posb--
	bita := (num & (1 << posa)) >> posa
	bitb := (num & (1 << posb)) >> posb
	return bita == bitb
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		num, bita, bitb := processLine(s.Text())
		fmt.Println(compareBits(num, uint(bita), uint(bitb)))
	}
}
