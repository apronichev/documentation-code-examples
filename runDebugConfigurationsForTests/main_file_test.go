package main

import (
	"testing"
)

func CalculateMulitplyTwo(x int) (result int) {
	return x * 2
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
func TestMultiply(t *testing.T) {
	if CalculateMulitplyTwo(0) != 2 {
		t.Error("Multiplication by zero must return zero")
	}
}
