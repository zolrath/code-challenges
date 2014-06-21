package main

import "fmt"

func main() {
	var t, n, a, b int

	// scan number of test cases
	fmt.Scanln("%d", &t)
	for i := 0; i < t; i++ {
		// scan number of vectors in set
		fmt.Scanln("%d", &n)
		for j := 0; j < n; j++ {
			// scan vector (a, b)
			fmt.Scanln("%d %d", &a, &b)
		}
	}
}
