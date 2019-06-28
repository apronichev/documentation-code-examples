package divisionMod

import "fmt"

func CalculateDivisionTwo(x int) (result int) {
	result = x / 2
	return result
}

func main() {
	result := CalculateDivisionTwo(4)
	fmt.Println(result)
}
