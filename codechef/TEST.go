package main

import "fmt"

func main() {
	var n int
	for {
		fmt.Scanf("%d", &n)
		if n == 42 {
			break
		}
		fmt.Println(n)
	}
}
