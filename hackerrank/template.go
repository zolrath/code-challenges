package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// Get test case from file, pack into Input struct.
func readNTestCases(s *bufio.Scanner, nCases int) []*Input {
	inputs := make([]*Input, nCases)
	for i := 0; i < nCases; i++ {
		s.Scan()
		line := readInts(s.Text())

		inputs[i] = &Input{
			Id: i,
			N:  line[0],
			C:  line[1],
			M:  line[2],
		}
	}
	return inputs
}

func readInts(line string) []int {
	lineSplit := strings.Split(line, " ")
	result := make([]int, len(lineSplit))
	for i, v := range lineSplit {
		result[i], _ = strconv.Atoi(v)
	}
	return result
}

func readFloats(line string) []float64 {
	lineSplit := strings.Split(line, " ")
	result := make([]float64, len(lineSplit))
	for i, v := range lineSplit {
		result[i], _ = strconv.ParseFloat(v, 64)
	}
	return result
}

// Contains the input required to solve a test case.
type Input struct {
	Id int
	N  int
	C  int
	M  int
}

// Contains the Id and Answer for a solved test case.
type Result struct {
	Id     int
	Answer int
}

// Calls the function which solves the test case.
func solver(in *Input, res chan<- *Result) {
	answer := countChocolate(in.N, in.C, in.M)
	res <- &Result{in.Id, answer}
}

func countChocolate(n, c, m int) int {
	// N = How much money Bob has.
	// C = Price of each chocolate.
	// M = How many wrappers it takes to get a free chocolate.
	wrappers := 0
	chocolates := 0
	for n >= c {
		n -= c
		chocolates++
		wrappers++
	}

	for wrappers >= m {
		wrappers -= m
		chocolates++
		wrappers++
	}
	return chocolates
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Get number of test cases
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	nCases, _ := strconv.Atoi(s.Text())

	// Read in n test cases, create slice of inputs.
	inputs := readNTestCases(s, nCases)
	results := collectResults(inputs, nCases)

	// Output results.
	for _, result := range results {
		fmt.Printf("%d\n", result.Answer)
	}
}

// collectResults sends each input to its own goroutine and collects the results
func collectResults(inputs []*Input, nCases int) []*Result {
	// Create result channel for go routines to send to.
	resultChan := make(chan *Result, nCases)

	// Create a new goroutine for each test case.
	spawnSolvers(inputs, resultChan)

	// Collect each test case result into a slice.
	results := make([]*Result, nCases)
	for i := 0; i < nCases; i++ {
		result := <-resultChan
		results[result.Id] = result
	}
	return results
}

// spawnSolvers creates a goroutine for each input, passing them the result channel.
func spawnSolvers(inputs []*Input, res chan<- *Result) {
	for _, input := range inputs {
		go solver(input, res)
	}
}
