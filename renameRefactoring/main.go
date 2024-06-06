package main

import (
	"errors"
	"fmt"
)

// findMaximumNumber returns the maximum number in a slice. If the slice is empty, it returns an error.
func findMaximumNumber(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("cannot find the maximum of an empty slice")
	}

	max := numbers[0]
	for _, num := range numbers[1:] {
		if num > max {
			max = num
		}
	}

	return max, nil
}

func main() {
	numbers := []int{3, 5, 7, 2, 9, 1}
	maxNumber, err := findMaximumNumber(numbers)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The maximum number is:", maxNumber)
	}
}
