package main

func Solve(arr []int) {
	result := arr[0]
	arr2 := arr
	for i := 0; i <= len(arr);  {
		for _, value := range arr2 {
			if arr[i] == value*-1 {
				arr = append(arr[:i], arr[i+1:]...)
				arr = append(arr[:i], arr[i+1:]...)
			}

		}
	}
	println(result)
}

func main() {
	Solve([]int{1, -1, 2, -2, 3})
	Solve([]int{-3, 1, 2, 3, -1, -4, -2})
	Solve([]int{1, -1, 2, -2, 3, 3})
	Solve([]int{-110, 110, -38, -38, -62, 62, -38, -38, -38})
	Solve([]int{-9, -105, -9, -9, -9, -9, 105})
}
