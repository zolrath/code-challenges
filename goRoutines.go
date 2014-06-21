package main

import (
	"fmt"
	"runtime"
)

type Input struct {
	Id int
	S  int // Starting number for sequence
}

type Result struct {
	Id     int
	Answer int
}

func solver(in *Input, r chan<- *Result) {
	answer := collatz(in.S)
	r <- &Result{in.Id, answer}
}

func collatz(n int) int {
	cycles := 1
	for i := n; i != 1; {
		switch i%2 == 0 {
		case true:
			i = i / 2
			cycles++
		case false:
			i = 3*i + 1
			cycles++
		}
	}
	return cycles
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	nCases := 100
	inputs := make([]*Input, nCases)
	resultChan := make(chan *Result, nCases)
	results := make([]int, nCases)

	for i := 0; i < nCases; i++ {
		inputs[i] = &Input{i, i + 1}
	}

	spawnSolvers(inputs, resultChan)

	for i := 0; i < nCases; i++ {
		res := <-resultChan
		results[res.Id] = res.Answer
	}

	for index, answer := range results {
		fmt.Printf("Case #%d: %d\n", index+1, answer)
	}
}

func spawnSolvers(inputs []*Input, r chan<- *Result) {
	for _, input := range inputs {
		go solver(input, r)
	}
}
