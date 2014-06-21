package main

import "fmt"

func printSlice(slice []int) {
	for i, v := range slice {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%d", v)
	}
	fmt.Print("\n")
}

func insertionSort(ar []int) []int {
	for i := 0; i < len(ar)-1; i++ {
		for j := i; j >= 0; j-- {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
		printSlice(ar)
	}
	return ar
}

func main() {
	var s, e int
	// get size of array
	fmt.Scanf("%d", &s)
	ar := make([]int, s)
	for i := 0; i < s; i++ {
		// get sorted array with one extra number
		fmt.Scanf("%d", &e)
		ar[i] = e
	}
	insertionSort(ar)
}
