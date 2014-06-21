package main

import "fmt"

func main() {
	var n, c int
	var cakes []int
	fmt.Scanf("%d", &n)
	cakes = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &c)
		cakes[i] = c
	}
	for _, v := range cakes {
		fmt.Println(v/2 + 1)
	}
}
