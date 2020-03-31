package main

import (
	"fmt"
	"os"
	"strconv"
)

func sum() {
	// An integer to copy: 95
	var sum int
	n, err := strconv.ParseInt(os.Args[1], 0, 0)

	if err != nil {
		fmt.Println("First input parameter must be integer")
		os.Exit(1)
	}

	for i := 0; i < int(n); i++ {
		sum += i
	}
	fmt.Print("Sum : ", sum, "\n")
}

func main() {
	sum()
}
