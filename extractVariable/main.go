package main

import "fmt"

// CalculateAverage calculates the average of a list of numbers.
func CalculateAverage(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	var total float64
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}

func main() {
	mySlice := []float64{10, 20, 30, 40, 50}
	avg := CalculateAverage(mySlice...)
	fmt.Printf("Average: %.2f\n", avg)
}
