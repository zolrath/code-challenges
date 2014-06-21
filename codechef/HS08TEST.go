package main

import "fmt"

func main() {
	var w int
	var b float32
	fmt.Scanf("%d %f", &w, &b)
	if w%5 == 0 && float32(w) <= (b+0.50) {
		b = b - float32(w) - 0.50
	}
	fmt.Printf("%1.2f\n", b)
}
