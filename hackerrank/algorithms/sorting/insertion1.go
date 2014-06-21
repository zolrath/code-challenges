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
	var v = ar[len(ar)-1]
	for i := len(ar) - 2; i >= 0; i-- {
		if ar[i] >= v {
			ar[i+1] = ar[i]
			printslice(ar)
		}
		if ar[i] < v {
			ar[i+1] = v
			printslice(ar)
			break
		}
		if i == 0 {
			ar[i] = v
			printslice(ar)
			break
		}
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
