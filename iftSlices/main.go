package main

import "fmt"

// Calculates the sum of all the numbers in a slice and passes to the printSum function
func sumNumbers(numbers ...int) {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	printSum(sum)
}

func printSum(s int)  {
	fmt.Println(s)
}

func main() {
	numbers := []int{5, 3, 4, 2, 5, 2, 62}
	sumNumbers(numbers)
}
