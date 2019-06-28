package main

import (
	"fmt"
)

func CalculatePlusTwo(x int) (result int) {
	result = x + 2
	return result
}

func main() {
	result := CalculatePlusTwo(4)
	fmt.Println(result)
}
