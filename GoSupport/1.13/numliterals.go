package main

import "fmt"

func shiftDemo(shift int) {
	x := 42
	result := x << shift
	fmt.Printf("shifting %d by %d results in %d\n", x, shift, result)
}

func main() {
	shiftDemo(0)
	shiftDemo(1)
}
