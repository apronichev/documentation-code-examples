package main

import (
	"fmt"
	"testing"
)

func CalculateMulitplyTwo(x int) (result int) {
	result = x * 2
	return result
}

func main() {
	result := CalculateMulitplyTwo(4)
	fmt.Println(result)
}

func TestMultiplyTable(t *testing.T) {
	tests := []struct {
		input      int
		wantResult int
	}{
		{2, 4},
		{6, 12},
		{-2, -4},
		{9999, 19998},
		{3, 6},
	}
	for _, tt := range tests {
		if gotResult := CalculateMulitplyTwo(tt.input); gotResult != tt.wantResult {
			t.Errorf("CalculateMulitplyTwo(%v) = %v, want %v", tt.input, gotResult, tt.wantResult)
		}
	}
}
func Testmultiply(t *testing.T) {
	if CalculateMulitplyTwo(0) != 0 {
		t.Error("Failed")
	}
}
