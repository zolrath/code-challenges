package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open(os.Args[1])
	i, _ := f.Stat()
	fmt.Println(i.Size())
}
